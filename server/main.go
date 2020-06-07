package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Panic("could not resolve server address")
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Panic("could not perform TCP listen")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	err := conn.SetReadDeadline(time.Now().Add(time.Millisecond * 20))
	if err != nil {
		log.Panic("could not set read deadline")
	}

	var buf bytes.Buffer
	io.Copy(&buf, conn)

	data := make([]byte, buf.Len())
	_, err = buf.Read(data)
	if err != nil {
		log.Panic("could not read buffer data")
	}
	fmt.Printf("\nData received: %s", data)

	daytime := time.Now().String()
	_, err = conn.Write([]byte(daytime))
	if err != nil {
		log.Panic("could not write response to TCP connection")
	}
}
