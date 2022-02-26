package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s, i)
	}

}

func main() {
	go f("-----------")
	go f("***********")
	time.Sleep(5 * time.Second)
}
