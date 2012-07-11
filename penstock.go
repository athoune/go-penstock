package main

import (
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
	err = client.Write(header, []byte("hello"))
	if err != nil {
		log.Panic(err)
	}
}
