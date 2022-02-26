package main

import "fmt"

func hello() {
	fmt.Println("hello")
}
func plus(val1 int, val2 int) int {
	return val1 + val2
}

func main() {
	hello()
	fmt.Println("result", plus(2, 3))
}
