package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi

}
func (r rectangle) perimeter() float64 {
	return (r.width + r.height) * 2
}
func (c circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func measure(s shape) {
	fmt.Println(s)
	fmt.Println("area", s.area())
	fmt.Println("perimeter", s.perimeter())
}

func main() {
	r := rectangle{width: 5, height: 10}
	c := circle{radius: 15}
	measure(r)
	measure(c)

}
