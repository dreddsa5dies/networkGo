package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	conn, _ := ln.Accept()

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print("Message Received:", string(msg))

		newMsg := strings.ToUpper(msg)

		conn.Write([]byte(newMsg + "\n"))
	}
}
