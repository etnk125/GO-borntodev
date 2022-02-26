package main

import "fmt"

type employee struct {
	id    string
	name  string
	phone string
}

func main() {
	employee1 := employee{
		id:    "101",
		name:  "pray",
		phone: "0900019283",
	}
	employees := [3]employee{employee1}
	employees[2] = employee{
		id:    "102",
		name:  "prayut",
		phone: "0900019284"}
	employees[1] = employee1
	fmt.Println("em :", employees)
	emslice := []employee{employee1}
	emslice = append(emslice, employees[2])
	fmt.Println("emslice", emslice)

}
