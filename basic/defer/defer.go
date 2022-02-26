package main

import "fmt"

func add(val1, val2 float64) {
	result := val1 + val2
	fmt.Println("result", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
	}
}
func dloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("d", i)
	}
}

func main() {
	fmt.Println("start")
	dloop()
	loop()
	fmt.Println("after loop ")
	defer fmt.Println("end")
	defer add(10, 100)
	defer add(20, 100)
	defer add(30, 100)

}
