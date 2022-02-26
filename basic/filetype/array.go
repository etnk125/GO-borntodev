package main

import "fmt"

var name [4]string

func main() {
	name[0] = "A"
	name[1] = "B"
	name[2] = "D"
	name[3] = "C"
	price := [4]float64{1, 3, 2, 4}
	fmt.Println(name)
	fmt.Println(price)
}
