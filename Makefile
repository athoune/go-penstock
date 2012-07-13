#include $(GOROOT)/src/Make.dist

all: client server

fmt:
	go fmt

protobuf:
	protoc --go_out=. *.proto

client: protobuf fmt
	go build penstock.go client.go header.pb.go

server: protobuf fmt
	go build penstockd.go server.go header.pb.go handler.go
