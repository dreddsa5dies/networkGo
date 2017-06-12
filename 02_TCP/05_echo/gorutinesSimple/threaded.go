/* ThreadedIPEchoServer */

package main

import (
	"log"
	"net"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveTCPAddr %v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("Ошибка ListenTCP %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		log.Print(conn.RemoteAddr(), " print: ", string(buf[0:]))

		_, err2 := conn.Write(buf[0:n])

		if err2 != nil {
			return
		}
	}
}
