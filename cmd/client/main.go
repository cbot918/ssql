package main

import (
	"fmt"
	"log"
	"net"
)

const (
	network = "tcp"
	port    = ":8888"
)

func main() {

	conn, err := net.Dial(network, port)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("hihi"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))

}
