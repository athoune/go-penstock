
fmt:
	go fmt

client: fmt
	go build penstock.go client.go

server: fmt
	go build penstockd.go server.go

all: client server

clean:
	rm penstock penstockd

