package main

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"net"
)

type client struct {
	conn net.Conn
}

func NewClient(host string, port int) (c client, err error) {
	conn, err := net.Dial("tcp", "localhost:4807") //[FIXME] build address
	if err != nil {
		return client{nil}, err
	}
	return client{conn}, nil
}

func (self *client) Write(header *Header, body []byte) error {
	header.Length = proto.Int32(int32(len(body)))
	var err error
	target, err := proto.Marshal(header)
	if err != nil {
		return err
	}
	err = binary.Write(self.conn, binary.LittleEndian, (int32)(len(target)))
	if err != nil {
		return err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		return err
	}
	_, err = self.conn.Write(body)
	if err != nil {
		return err
	}
	return nil
}

type writer interface {
	Write(b []byte) (n int, err error)
}

func (self *client) NewWriter(header *Header) (w writer, er error) {
	var err error
	target, err := proto.Marshal(header)
	if err != nil {
		return nil, err
	}
	err = binary.Write(self.conn, binary.LittleEndian, len(target))
	if err != nil {
		return nil, err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		return nil, err
	}
	// [FIXME] the writer must whining when someone try to write more data than announced
	return self.conn, nil
}
