package main

import (
	"io"
	"log"
)

type Handler interface {
	Handle(header *Header, body io.Reader)
}

type DebugHandler struct {
}

func (self *DebugHandler) Handle(header *Header, body io.Reader) {
	log.Println("header:", header)
	data := make([]byte, header.GetLength())
	_, err := body.Read(data)
	if err != nil {
		log.Println("body error:", err)
	}
	log.Println("body:", data)
}
