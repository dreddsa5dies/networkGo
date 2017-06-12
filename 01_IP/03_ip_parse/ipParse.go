/*	IP	*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s IP\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	// проверка на возможность IP
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("Айпишник ", addr.String())
	}
	os.Exit(0)
}
