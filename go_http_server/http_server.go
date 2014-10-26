// http_go.go
package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"-"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
	memo    string
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}
