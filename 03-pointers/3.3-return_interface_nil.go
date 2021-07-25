package main

import (
	"fmt"
	"net/http"
)

// custom interface
type newError interface {
	GetSummary() string
}

// our struct which implements both error interface and custom interface
type customError struct {
	status  int
	message string
}

// implementation
func (c customError) Error() string {
	return c.message
}

// implementation
func (c customError) GetSummary() string {
	return fmt.Sprintf("returned '%v' with status code of %d", c.message, c.status)
}

type person struct {
	name string
	age  int
}

func (p person) save() (*person, newError) {
	if p.name != "foo" {
		return nil, customError{http.StatusNotFound, "No resource found"}
	}
	return &person{p.name, p.age}, nil
}

func main() {
	fmt.Println("return interface, return object")
	p1 := person{"bar", 23}
	savedPerson1, err := p1.save()
	if err != nil {
		fmt.Println("p1 could not be saved !")
	}
	if savedPerson1 == nil {
		fmt.Println("Expected person NOT to be saved")
	}

	fmt.Println("return interface, return nil")
	p2 := person{"foo", 23}
	savedPerson2, err := p2.save()
	if err == nil {
		fmt.Println("expected error be nil")
		fmt.Println("saved person details, Name:", savedPerson2.name, ", age:", savedPerson2.age)
	}
}
