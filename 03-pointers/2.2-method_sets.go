package main

import (
	"fmt"
	"math"
)

// *************************************

type shape interface {
	area() float64
}

// *************************************
// 'circle' implements 'shape'
type circle struct {
	radius float64
}

func (c *circle) area() float64 { // pointer receiver
	return math.Pi * c.radius * c.radius
}

// *************************************

func info(s shape) {
	fmt.Println("area", s.area())
}

// *************************************

func main() {
	c := circle{5}
	info(&c) // pointer value
}
