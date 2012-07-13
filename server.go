package main

import (
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"io"
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
	var size int32
	err := binary.Read(conn, binary.LittleEndian, &size)
	if err != nil {
		log.Println(err)
	}
	header := &Header{}
	data := make([]byte, size)
	_, err = conn.Read(data)
	if err != nil {
		log.Println(err)
	}
	err = proto.Unmarshal(data, header)
	if err != nil {
		log.Println(err)
	}
	handler := new(DebugHandler)
	handler.Handle(header, io.LimitReader(conn, int64(header.GetLength())))
}
