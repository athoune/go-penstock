package main

import (
	"encoding/binary"
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
	var size int
	err := binary.Read(conn, binary.LittleEndian, &size)
	log.Println(size)
	target := make([]byte, size)
	_, err = conn.Read(target)
	err = binary.Read(conn, binary.LittleEndian, &size)
	log.Println(size)
	body := make([]byte, size)
	_, err = conn.Read(body)
	if err != nil {
		log.Println(err)
	}
	log.Println(body)
}
