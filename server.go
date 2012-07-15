package main

import (
	"log"
	"net"
)

type server struct {
	conn net.Listener
}

func NewServer(port int) (s server, err error) {
	ln, err := net.Listen("tcp", ":4807")
	if err != nil {
		return server{nil}, err
	}
	log.Println("Starting the server")
	// [FIXME] handles loop and listen a chan for stopping.
	for {
		conn, err := ln.Accept()
		if err != nil {
			//error
			continue
		}
		log.Println("Handling a connection")
		go handleConnection(conn)
	}
	return server{ln}, nil
}

func handleConnection(conn net.Conn) {
	handler := new(DebugHandler)
	message, err := ReadMessage(conn)
	if err != nil {
		log.Println(err)
	}
	handler.Handle(message)
}
