/* LoadGob */

package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
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
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person

	loadGob("./person.gob", &person)

	fmt.Println("Person", person.String())
}

func loadGob(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошибка os.Open %v", err)
	}

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	if err != nil {
		log.Fatalf("Ошибка decoder.Decode %v", err)
	}

	inFile.Close()
}
