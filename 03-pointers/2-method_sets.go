/*
	Method sets: 'methods' attached to a 'type' are called 'method sets'
	1) receiver can be of pointer or non-pointer type
	2) receiver can be of pointer only

	Receivers       Values
	-----------------------------------------------
	(t T)           T and *T
	(t *T)          *T
	(t *T)			T 		// DOES NOT WORK !!

*/

package main

import (
	"fmt"
	"math"
)

// *************************************
// interface

type shape interface {
	area() float64
}

// *************************************
// one who is implementing the interface

type circle struct {
	radius float64
}

// *************************************
// method which is attached to 'circle' model

func (c circle) area() float64 { // RECEIVER
	return math.Pi * c.radius * c.radius
}

// *************************************
// utility methhos to calculate the area

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(c) // VALUES
}
