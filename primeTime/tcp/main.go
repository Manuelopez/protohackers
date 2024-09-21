package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
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
	curr := ""

	for {
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("ERROR reading")
			}
			return
		}

		curr += string(buf[:n])

		fmt.Println("current json formed: ", curr)
		t := strings.Contains(curr, "\n")
		if t {

			strs := strings.Split(curr, "\n")
			req := strs[0]

			fmt.Println("complet json formed", curr, "the actual", req)
			if len(strs) > 1 {
				curr = strings.Join(strs[1:], "")

				fmt.Println("length was more than one")
			} else {
				curr = ""

				fmt.Println("length was ONLY one")
			}
			data := op{}
			err := json.Unmarshal([]byte(req), &data)
			if err != nil {
				c.Write([]byte("{\"method\":\"isPrime\",\"prime\":123}"))
				fmt.Println("error ddecoding json")
			} else {

				fmt.Println("json decoded")
				if big.NewInt(int64(data.Number)).ProbablyPrime(0) {
					c.Write([]byte("{\"method\":\"isPrime\",\"prime\":true}"))
				} else {
					c.Write([]byte("{\"method\":\"isPrime\",\"prime\":false}"))
				}
			}
		}
	}
}
