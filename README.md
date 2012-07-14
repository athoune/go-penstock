High pressure communication for Go
==================================

Penstock is a low level tool to build event oriented communication in
a connected socket. It borrows ideas from XMPP's IQ, JSON-RPC, SPDY, HTTP.

Penstock handles flow of bytes, so, it can be used to transfer files or
use any serialized values.

Protobuf is used for headers, but you can use any serialization
(json, netstrings, msgpack, xml…) to communicate.

Features
--------

 * √ One persistant connection
 * √ Header
 * √ Read and write body as a flow
 * √ Server handlers
 * _ Registering different server handlers
 * _ Errors
 * _ Body compression
 * _ Body checksum
 * _ Throttling
 * _ File transfert example
 * _ Protobuf RPC
 * _ Client example in another language (python, nodejs…)

Try it
------

penstock.go is a client example, and penstockd.go is a server example.

Modify it
---------

You should install [goprotobuf](https://code.google.com/p/goprotobuf/).

A Makefile is used to handle compilation.

    make all


Licence
-------

BSD.
