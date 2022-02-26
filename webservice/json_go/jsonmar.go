package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {
	data, _ := json.Marshal(&employee{101, "py chan", "12123131312", "prewit@mail.com"})
	fmt.Println(string(data))
}
