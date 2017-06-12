/* SaveJSON */

package main

import (
	"encoding/json"
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

	saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Ошибка os.Create %v", err)
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)

	err = encoder.Encode(key)
	if err != nil {
		log.Fatalf("Ошибка encoder.Encode %v", err)
	}
}
