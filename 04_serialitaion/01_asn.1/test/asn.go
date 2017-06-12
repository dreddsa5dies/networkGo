// Abstract Syntax Notation One
// ASN.1

package main

import (
	"encoding/asn1"
	"fmt"
	"log"
)

func main() {
	// int
	mdata, err := asn1.Marshal(13)
	if err != nil {
		log.Fatalf("Ошибка Marshal %v", err)
	}
	fmt.Println("До marshal/unmarshal: ", mdata)

	var n int
	_, err = asn1.Unmarshal(mdata, &n)
	if err != nil {
		log.Fatalf("Ошибка Unmarshal %v", err)
	}

	fmt.Println("После marshal/unmarshal: ", n)

	// строка
	s := "hello"
	mdata, err = asn1.Marshal(s)
	if err != nil {
		log.Fatalf("Ошибка Marshal %v", err)
	}
	fmt.Println("До marshal/unmarshal: ", mdata)
	var newstr string
	asn1.Unmarshal(mdata, &newstr)
	fmt.Println("После marshal/unmarshal: ", newstr)
}
