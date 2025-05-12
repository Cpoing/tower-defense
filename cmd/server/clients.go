package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// make sure only the producer sends packets

func clients() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Unable to connect to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var frame strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected:", err)
			break
		}

		if strings.TrimSpace(line) == "<<<END>>>" {
			fmt.Print("\033[H\033[2J")
			fmt.Print(frame.String())
			frame.Reset()
		} else {
			frame.WriteString(line)
		}
	}
}

