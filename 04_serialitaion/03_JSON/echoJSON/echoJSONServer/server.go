/* JSON	EchoServer */

package main

import (
	"encoding/json"
	"log"
	"net"
)

// Person struct
type Person struct {
	Name  Name
	Email []Email
}

// Name struct
type Name struct {
	Family   string
	Personal string
}

// Email struct
type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":	" + v.Address
	}
	return s
}

func main() {
	service := "0.0.0.0:1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
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

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		var person Person
		decoder.Decode(&person)
		encoder.Encode(person)

		log.Printf("Connect ok %v", conn.RemoteAddr())
		log.Printf("Send data \n%v", person)

		conn.Close() //	we're finished
	}
}
