package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	fmt.Println("started")
	listener, err := net.Listen("tcp", ":7")
	if err != nil {
		fmt.Println("ERROR listener")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR accept")
			return
		}

		go handleConnection(conn)

	}

}

type op struct {
	Method string `json:"method"`
	Number int    `json:"number"`
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


    curr := string(buf[:n])
		fmt.Println("current json formed: ", curr)
		t := strings.Contains(curr, "\n")
		if t {


			data := op{}
			err := json.Unmarshal([]byte(curr), &data)
			if err != nil {
				c.Write([]byte("{\"method\":\"isPrime\",\"prime\":123}"))
				fmt.Println("error ddecoding json")
				return
			} else {

				fmt.Println("json decoded")
				if isPrime(data.Number) {

					fmt.Println("json true")
					c.Write([]byte("{\"method\":\"isPrime\",\"prime\":true}"))

					fmt.Println("did it send the write???")
				} else {

					fmt.Println("json false")
					c.Write([]byte("{\"method\":\"isPrime\",\"prime\":false}"))

					fmt.Println("did it send the write???")
				}
			}
		}
	}
}

func isPrime(n int) bool {
  if n < 0{return false}
	if n >= 0 && n <= 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
