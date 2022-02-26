package main

import (
	"fmt"
)

func dumbgrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func setMinMax(score int) int {
	if score > 100 {
		return 100
	}
	if score < 0 {
		return 0
	}
	return score
}

func easygrade(score int) string {
	return string("FFFFFDCBAAA"[int(setMinMax(score)/10)])
}

func main() {
	for score := 0; score <= 100; score += 10 {
		fmt.Println(score)
		fmt.Println(dumbgrade(score))
		fmt.Println(easygrade(score))
	}
}
