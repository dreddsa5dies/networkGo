/* UDPDaytimeServer */

package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveUDPAddr %v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Ошибка ListenUDP %v", err)
	}

	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	log.Print("connect ", addr.IP.String(), ":", addr.Port)
	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}
