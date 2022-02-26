package main

import "fmt"

var product = make(map[string]float64)

func main() {
	//init
	fmt.Println("init", product)
	//add
	product["mc"] = 400
	product["mc2"] = 300.2
	product["mc4"] = 4100.9999
	fmt.Println("added", product)
	//delete
	delete(product, "mc")
	fmt.Println("deleted", product)
	//update
	product["mc4"] = 410
	fmt.Println("product", product)

	//access
	fmt.Println("access mc4", product["mc4"])

	//short hand
	shortMap := map[string]string{"1": "one", "2": "two"}
	fmt.Println("short hand map", shortMap)
}
