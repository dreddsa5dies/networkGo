/* SaveGob */

package main

import (
	"encoding/gob"
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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveGob("./person.gob", person)
}

func saveGob(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Ошибка os.Create %v", err)
	}

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	if err != nil {
		log.Fatalf("Ошибка gob.NewEncoder %v", err)
	}

	outFile.Close()
}
