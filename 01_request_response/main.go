package main

import (
	"fmt"
	"log"
	"net"
)

func respond(conn net.Conn) {
	defer conn.Close()

	hello := "<!DOCTYPE html><head><title>Hello World!</title></head><body><h1>Hello World!</h1></body></html>"

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(hello))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, hello)
}

func main() {
	app, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	for {
		conn, err := app.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go respond(conn)
	}
}
