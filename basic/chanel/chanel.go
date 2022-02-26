// communicate between goroutine
// thread in java

package main

import (
	"fmt"
	"time"
)

func process(ch chan string, data string) {
	ch <- data
}

func main() {
	ch := make(chan string)
	go process(ch, "new data")
	fmt.Println(<-ch)
	time.Sleep(100 * time.Millisecond)
}
