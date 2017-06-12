/* ASN1 DaytimeServer */

package main

import (
	"encoding/asn1"
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
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

		daytime := time.Now()
		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)
		log.Print("Data ", mdata)
		conn.Write(mdata)
		conn.Close() // we're finished
		log.Print("Data send to ", conn.RemoteAddr())
	}
}
