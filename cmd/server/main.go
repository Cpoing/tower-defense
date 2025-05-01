package main

import (
	"fmt"
	"net"

	"tower-defense/pkg/connection"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting tcp server")
		return
	}
	defer ln.Close()
	fmt.Println("TCP Server Started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error connecting to server")
			return
		}

		go connection.HandleConnection(conn)
	}
}
