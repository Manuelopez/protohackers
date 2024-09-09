package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:200")
	if err != nil {
		fmt.Println("ERROR listener")
		return
	}

	for i := 0; i < 5; i++ {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR accept")
			return

		}

		go handleConnection(conn)

	}

}

func handleConnection(c net.Conn) {

	defer func() { c.Close() }()

	buf := make([]byte, 1024)

	for {
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("ERROR reading")
			}
      return
		}
    c.Write(buf[:n])

	}
}
