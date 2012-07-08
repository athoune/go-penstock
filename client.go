package main

import (
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

func (self *client) Write(target []byte, body []byte) error {
	var err error
	err = binary.Write(self.conn, binary.LittleEndian, len(target))
	if err != nil {
		return err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		return err
	}
	err = binary.Write(self.conn, binary.LittleEndian, len(body))
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

func (self *client) NewWriter(target []byte, size int) (w writer, er error) {
	var err error
	err = binary.Write(self.conn, binary.LittleEndian, len(target))
	if err != nil {
		return nil, err
	}
	_, err = self.conn.Write(target)
	if err != nil {
		return nil, err
	}
	err = binary.Write(self.conn, binary.LittleEndian, size)
	if err != nil {
		return nil, err
	}
	// [FIXME] the writer must whining when someone try to write more data than announced
	return self.conn, nil
}
