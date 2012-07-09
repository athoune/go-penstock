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
	err = client.Write([]byte(""), []byte("hello"))
	if err != nil {
		log.Panic(err)
	}
}
