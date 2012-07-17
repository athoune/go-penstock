package main

import (
	"log"
	"net"
)

// Each read messages are handled by a Handler.
// Be careful, if you don't read the body and try to fetch a new message,
// it will crash.
type Handler interface {
	Handle(message *Message)
}

// Handler for displaying message.
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

// Handler wich send an ok response.
type AckHandler struct {
	conn net.Conn
}

func (self AckHandler) Handle(message *Message) {
	data := make([]byte, message.Header.GetLength())
	_, _ = message.Body.Read(data)
	log.Println("body:", data)
	//[FIXME] check message.Header.Type == Header_QUERY
	r := Header_RESPONSE
	header := &Header{
		Id:   message.Header.Id,
		Path: message.Header.GetPath(),
		Type: &r,
	}
	WriteMessage(self.conn, NewBytesMessage(header, []byte("ok")))
}

type CompleteMessage struct {
	Header *Header
	Body   []byte
}

func NewCompleteMessage(source *Message) (complete *CompleteMessage, err error) {
	data := make([]byte, source.Header.GetLength())
	_, err = source.Body.Read(data)
	return &CompleteMessage{source.Header, data}, nil
}

type ChanHandler struct {
	response chan CompleteMessage
}

func (self ChanHandler) Handle(message *Message) {
	complete, _ := NewCompleteMessage(message)
	self.response <- *complete
}
