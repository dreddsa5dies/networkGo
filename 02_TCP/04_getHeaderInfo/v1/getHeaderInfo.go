/*	GetHeadInfo */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveTCPAddr %v", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Ошибка DialTCP %v", err)
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatalf("Ошибка conn.Write %v", err)
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalf("Ошибка ioutil.ReadAll %v", err)
	}

	fmt.Println(string(result))
	os.Exit(0)
}
