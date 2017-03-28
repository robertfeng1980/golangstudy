package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Printf(" 这是客户端  \n")
	conn, err := net.Dial("tcp", "192.168.23.196:8888")

	checkError(err)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "hello server\n")
	data, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		panic(err)
	}

	fmt.Printf("  ===========   ", data)

	conn.Close()

}
