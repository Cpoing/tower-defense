package connection

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	fmt.Println(conn)
	fmt.Println("Connected to server")
}
