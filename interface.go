package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) area() float64 {
	return r.width * r.height
}

type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interface example")

	c1 := circle{4.5}
	r1 := rect{5, 7}
	fmt.Println(c1.area(), r1.area())

	shapes := []shape{c1, r1}
	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
