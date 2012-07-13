package main

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"io"
	"net"
)

type client struct {
	current_id uint32
	conn       net.Conn
}

func NewClient(host string, port int) (c *client, err error) {
	conn, err := net.Dial("tcp", "localhost:4807") //[FIXME] build address
	if err != nil {
		return &client{0, nil}, err
	}
	return &client{0, conn}, nil
}

func (self *client) Write(header *Header, body []byte) error {
	header.Length = proto.Int32(int32(len(body)))
	self.current_id += 1
	header.Id = proto.Uint32(self.current_id)
	var err error
	target, err := proto.Marshal(header)
	if err != nil {
		self.conn.Close()
		return err
	}
	err = binary.Write(self.conn, binary.LittleEndian, (int32)(len(target)))
	if err != nil {
		self.conn.Close()
		return err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		self.conn.Close()
		return err
	}
	_, err = self.conn.Write(body)
	if err != nil {
		self.conn.Close()
		return err
	}
	self.current_id += 1
	return nil
}

type smallWriter struct {
	max    int32
	size   int32
	writer io.Writer
}

func NewSmallWriter(size int32, writer io.Writer) *smallWriter {
	return &smallWriter{size, 0, writer}
}

func (self *smallWriter) Write(p []byte) (n int, err error) {
	if self.size+int32(len(p)) > self.max {
		return int(self.size), *new(error) // "Too many stuff written")
	}
	self.size += int32(len(p))
	return self.writer.Write(p)
}

func (self *client) NewWriter(header *Header) (w io.Writer, er error) {
	var err error
	target, err := proto.Marshal(header)
	if err != nil {
		self.conn.Close()
		return nil, err
	}
	err = binary.Write(self.conn, binary.LittleEndian, len(target))
	if err != nil {
		self.conn.Close()
		return nil, err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		self.conn.Close()
		return nil, err
	}
	return NewSmallWriter(header.GetLength(), self.conn), nil
}
