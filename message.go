package main

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"io"
)

type Message struct {
	*Header
	Body io.Reader
}

func ReadMessage(conn io.Reader) (message Message, err error) {
	var size int32
	err = binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		return Message{}, err
	}
	header := &Header{}
	data := make([]byte, size)
	_, err = conn.Read(data)
	if err != nil {
		return Message{}, err
	}
	err = proto.Unmarshal(data, header)
	if err != nil {
		return Message{}, err
	}
	return Message{header, io.LimitReader(conn, int64(header.GetLength()))}, nil
}
