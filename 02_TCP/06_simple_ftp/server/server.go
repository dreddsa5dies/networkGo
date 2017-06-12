/* FTP Server */

package main

import (
	"log"
	"net"
	"os"
)

const (
	// DIR command
	DIR = "DIR"
	// CD command
	CD = "CD"
	// PWD command
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:1202"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveIPAddr %v", err)
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
			conn.Close()
			log.Printf("connect: %v FAILRUE", conn.RemoteAddr())
			return
		}

		log.Printf("connect: %v", conn.RemoteAddr())

		s := string(buf[0:n])

		// decode request
		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
			dirList(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
		log.Printf("%v OK %v", conn.RemoteAddr(), CD)
	} else {
		conn.Write([]byte("ERROR"))
		log.Printf("%v FAILRUE %v", conn.RemoteAddr(), CD)
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		log.Printf("%v FAILRUE %v", conn.RemoteAddr(), PWD)
		return
	}
	conn.Write([]byte(s))
	log.Printf("%v OK %v", conn.RemoteAddr(), PWD)
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		log.Printf("%v FAILRUE %v", conn.RemoteAddr(), DIR)
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		log.Printf("%v FAILRUE %v", conn.RemoteAddr(), DIR)
		return
	}

	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}

	log.Printf("%v OK %v", conn.RemoteAddr(), DIR)
}
