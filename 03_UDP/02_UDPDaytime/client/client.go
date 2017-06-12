/* UDPDaytimeClient */

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveUDPAddr %v", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Ошибка DialUDP %v", err)
	}

	_, err = conn.Write([]byte("hello, get time"))
	if err != nil {
		log.Fatalf("Ошибка conn.Write %v", err)
	}

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Fatalf("Ошибка ResolveTCPAddr %v", err)
	}

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
