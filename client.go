package main

import (
	"net"
)

type client struct {
	current_id uint32
	conn       net.Conn
	response   chan CompleteMessage
}

func NewClient(host string, port int) (c *client, err error) {
	conn, err := net.Dial("tcp", "localhost:4807") //[FIXME] build address
	if err != nil {
		return &client{0, nil, nil}, err
	}
	c = &client{0, conn, make(chan CompleteMessage)}
	response := c.response
	handler := ChanHandler{response}
	go ReadLoop(conn, handler)
	return
}

func (self *client) Write(message *Message) error {
	return WriteMessage(self.conn, message)
}
