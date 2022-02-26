package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// write
	data := []byte("hello\n eiei")
	err := os.WriteFile("data.txt", data, 0644)
	check(err)
	// create
	f, err := os.Create("names.txt")
	check(err)
	f.Close()

	data2 := []byte("name1\nnamw2")
	os.WriteFile("names.txt", data2, 0644)
}
