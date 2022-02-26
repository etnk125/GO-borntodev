package main

import "fmt"

var numberInt, num2 int = 100, 20

var msg string = "hi"

func main() {
	numberFloat := 25.4

	fmt.Println(numberInt)
	fmt.Println(numberFloat)
	fmt.Println(msg, "world")
	fmt.Println(float64(numberInt) + numberFloat)
	fmt.Println("myint =", num2)
}
