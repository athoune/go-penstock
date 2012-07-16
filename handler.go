package main

import (
	"log"
	"net"
)

type Handler interface {
	Handle(message *Message)
}

type DebugHandler struct {
}

func (self DebugHandler) Handle(message *Message) {
	log.Println("header:", message.Header)
	data := make([]byte, message.Header.GetLength())
	_, err := message.Body.Read(data)
	if err != nil {
		log.Println("body error:", err)
	}
	log.Println("body:", data)
}

type AckHandler struct {
	conn net.Conn
}

func (self AckHandler) Handle(message *Message) {
	/*data := make([]byte, message.Header.GetLength())*/
	/*_, err := message.Body.Read(data)*/
	header := &Header{Path: message.Header.GetPath()}
	WriteMessage(self.conn, NewBytesMessage(header, []byte("ok")))
}
