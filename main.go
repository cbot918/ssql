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

	fmt.Println("listening ", port)

	lis, err := net.Listen(network, port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(buf[:n]))

	_, err = conn.Write([]byte("reply"))
	if err != nil {
		fmt.Println(err)
	}

}
