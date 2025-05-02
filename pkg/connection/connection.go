package connection

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]bool)
	clientsMu sync.Mutex
	producer  net.Conn
	prodOnce  sync.Once
)

func HandleConnection(conn net.Conn) {
	var isProducer bool

	prodOnce.Do(func() {
		producer = conn
		isProducer = true
	})

	if isProducer {
		fmt.Println("Producer connected:", conn.RemoteAddr())
		handleProducer(conn)
	} else {
		fmt.Println("Client connected:", conn.RemoteAddr())
		addClient(conn)
	}
}

func handleProducer(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		data := scanner.Text() + "\n"
		broadcast(data)
	}
	fmt.Println("Producer disconnected")
	conn.Close()
}

func addClient(conn net.Conn) {
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()

	go func() {
		defer func() {
			clientsMu.Lock()
			delete(clients, conn)
			clientsMu.Unlock()
			conn.Close()
			fmt.Println("Client disconnected:", conn.RemoteAddr())
		}()
		buf := make([]byte, 1)
		for {
			_, err := conn.Read(buf)
			if err != nil {
				return
			}
		}
	}()
}

func broadcast(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		_, err := fmt.Fprint(conn, message)
		if err != nil {
			fmt.Println("Error sending to client:", conn.RemoteAddr())
			conn.Close()
			delete(clients, conn)
		}
	}
}
