package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(title string) float64 {
	fmt.Print(title)
	input, _ := reader.ReadString('\n')
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}
	return val

}

func getOperator() string {
	fmt.Print("[+ - * /] : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func operations(op string, num1 *float64, num2 float64) {
	switch op {
	case "+":
		*num1 += num2
	case "-":
		*num1 -= num2
	case "*":
		*num1 *= num2
	case "/":
		*num1 /= num2
	default:
		panic("invalid operator")
	}
}

func main() {
	round := int(getInput("round :"))
	money := getInput("money :")

	for i := 0; i < round-1; i++ {

		operations(getOperator(), &money, getInput("do with :"))
	}
	fmt.Println("money : ", money)
}
