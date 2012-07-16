package main

import (
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
	handler := DebugHandler{}
	go ReadLoop(conn, handler)
	return &client{0, conn}, nil
}

func (self *client) Write(message *Message) error {
	return WriteMessage(self.conn, message)
}
