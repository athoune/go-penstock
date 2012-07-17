package main

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Message struct {
	*Header
	Body io.Reader
}

func (self Message) Read() (body []byte, err error) {
	body = make([]byte, self.Header.GetLength())
	_, err = self.Body.Read(body)
	if err != nil {
		return nil, err
	}
	return
}

func NewBytesMessage(header *Header, body []byte) *Message {
	header.Length = proto.Int32(int32(len(body)))
	return &Message{header, bytes.NewBuffer(body)}
}

func ReadMessage(conn io.Reader) (message *Message, err error) {
	var size int32
	err = binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		return &Message{}, err
	}
	header := &Header{}
	data := make([]byte, size)
	_, err = conn.Read(data)
	if err != nil {
		return &Message{}, err
	}
	err = proto.Unmarshal(data, header)
	if err != nil {
		return &Message{}, err
	}
	return &Message{header, io.LimitReader(conn, int64(header.GetLength()))}, nil
}

type TooBigError string

func (e TooBigError) Error() string {
	return string(e)
}

// Copy the content of the reader. Be careful, only length announced in the header
// will be copied.
func WriteMessage(conn net.Conn, message *Message) error {
	length := int64(message.Header.GetLength())
	target, err := proto.Marshal(message.Header)
	if err != nil {
		conn.Close()
		return err
	}
	err = binary.Write(conn, binary.LittleEndian, int32(len(target)))
	if err != nil {
		conn.Close()
		return err
	}
	_, err = conn.Write(target)
	if err != nil {
		conn.Close()
		return err
	}
	r := io.LimitReader(message.Body, length)
	var real_size int64
	real_size, err = io.Copy(conn, r)
	if err != nil {
		conn.Close()
		return err
	}
	if real_size < length {
		return TooBigError("Length is bigger than reader capacity")
	}
	return nil

}

func ReadLoop(conn net.Conn, handler Handler) {
	loop := true
	for loop {
		message, err := ReadMessage(conn)
		if err != nil {
			if err != io.EOF {
				log.Println("Error while reading message", err)
			}
			loop = false
			conn.Close()
		} else {
			handler.Handle(message)
		}
	}
}
