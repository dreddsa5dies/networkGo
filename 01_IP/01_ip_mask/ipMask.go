/*	Mask */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Использование: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is \t\t", addr.String(),
		"\nDefault mask length is \t", bits,
		"\nLeading ones count is \t", ones,
		"\nMask is (hex) \t\t", mask.String(),
		"\nNetwork is \t\t", network.String())
	os.Exit(0)
}
