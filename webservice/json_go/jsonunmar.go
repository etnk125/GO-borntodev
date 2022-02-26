package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {
	aemployee := employee{}
	err := json.Unmarshal([]byte(`{"ID":101, "Name":"py chan","Tel": "12123131312", "email":"prewit@mail.com"}`), &aemployee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aemployee)
	fmt.Println(aemployee.Name)
}
