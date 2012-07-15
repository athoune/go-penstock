package main

import (
	"log"
)

type Handler interface {
	Handle(message Message)
}

type DebugHandler struct {
}

func (self *DebugHandler) Handle(message Message) {
	log.Println("header:", message.Header)
	data := make([]byte, message.Header.GetLength())
	_, err := message.Body.Read(data)
	if err != nil {
		log.Println("body error:", err)
	}
	log.Println("body:", data)
}
