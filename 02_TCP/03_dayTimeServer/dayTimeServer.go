/* DaytimeServer */

package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
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

		log.Print("Connect ok ", conn.RemoteAddr())

		daytime := time.Now().String() + "\n"
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
