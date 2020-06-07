package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port numberOfTCPCalls", os.Args[0])
		os.Exit(1)
	}

	callsArg := os.Args[2]
	numberOfCalls, err := strconv.Atoi(callsArg)
	if err != nil {
		log.Panic("could not parse convert numberOfTCPCalls")
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Panic("could not parse server address")
	}

	for i := 0; i < numberOfCalls; i++ {
		go doTCPCall(tcpAddr, service, i)
	}

	time.Sleep(time.Second)
	os.Exit(0)
}

func doTCPCall(tcpAddr *net.TCPAddr, service string, routine int) {
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("could create TCP connection from routine %d to server %s", routine, service)
	}

	data := "HEAD / HTTP/1.0 \r\n"
	fmt.Printf("Routine %d, writting %s to %s \n", routine, data, service)
	_, err = conn.Write([]byte(data))
	if err != nil {
		fmt.Printf("could not write data of routine %d into its TCP connection", routine)
	}
	_, err = conn.Write([]byte(fmt.Sprintf("Hello from Routine %d \n", routine)))
	if err != nil {
		fmt.Printf("could not write data of routine %d into its TCP connection", routine)
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil && err != io.EOF {
		log.Fatalf("could not read data of routine %d from its TCP connection", routine)
	}

	fmt.Printf("Routine %d, reading from TCP connection: %s \n", routine, string(result))
}
