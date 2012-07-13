include $(GOROOT)/src/Make.dist

fmt:
	go fmt

protobuf:
	protoc --go_out=. *.proto

client: protobuf fmt
	go build penstock.go client.go header.pb.go

server: protobuf fmt
	go build penstockd.go server.go header.pb.go handler.go

all: client server

include $(GOROOT)/src/pkg/code.google.com/p/goprotobuf/Make.protobuf

