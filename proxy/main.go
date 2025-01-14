// Proxy mediates TCP traffic between client and upstream. Adapted from
// youtu.be/J4J-A9tcjcA.
//
// Level: intermediate
// Topics: net, security, concurrency
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// NOTE: don't put any blocking code here!
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	// Release precious (there are not that many) file descriptor.
	defer conn.Close()

	upstream, err := net.Dial("tcp", "google.com:http")
	if err != nil {
		log.Print(err)
		return
	}
	defer upstream.Close()

	// In this case it's ok not track the goroutine.
	go io.Copy(upstream, conn)
	io.Copy(conn, upstream)
}
