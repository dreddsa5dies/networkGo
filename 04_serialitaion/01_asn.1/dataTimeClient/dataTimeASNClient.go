/* ASN.1 DaytimeClient */

package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	if err != nil {
		log.Fatalf("Ошибка net.Dial %v", err)
	}

	result, err := readFully(conn)
	if err != nil {
		log.Fatalf("Ошибка readFully %v", err)
	}

	var newtime time.Time
	_, err = asn1.Unmarshal(result, &newtime)
	if err != nil {
		log.Fatalf("Ошибка Unmarshal %v", err)
	}

	fmt.Println("После marshal/unmarshal: ", newtime.String())
	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
