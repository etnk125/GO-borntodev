package main

import "fmt"

func passpointer(ipoint *int) {
	*ipoint = 0
}

func main() {
	i := 1
	fmt.Println("i", i)
	passpointer(&i)
	fmt.Println("i", i)
}
