package main

import (
	"bytes"
	"log"
)

func main() {
	client, err := NewClient("localhost", 4807)
	if err != nil {
		//error
	}
	// some login?
	header := &Header{
		Path: []byte("some/path"),
	}
	log.Println(header)
	err = client.Write(NewBytesMessage(header, []byte("hello")))
	if err != nil {
		log.Panic(err)
	}
	err = client.Write(NewBytesMessage(header, []byte("world!")))
	if err != nil {
		log.Panic(err)
	}
	var msg CompleteMessage
	for {
		msg = <-client.response
		b := bytes.NewBuffer(msg.Body)
		log.Println("Response:", msg.Header, b.String())
	}
}
