package main

import "fmt"

func main() {
	var name []string
	name = []string{"a", "b", "c"}
	fmt.Println(name)
	name = append(name, "d", "E", "F#", "G")
	fmt.Println(name)
	name2 := name[4:7]
	fmt.Println(name2)
	name2 = name[0:4]
	fmt.Println(name2)
}
