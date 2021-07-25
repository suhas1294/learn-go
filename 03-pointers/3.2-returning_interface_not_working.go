package main

import "fmt"

type human interface {
	speak() string
	clone() person
}

type person struct {
	name string
	age  int
}

func (p person) speak() string {
	return fmt.Sprintf("My name is %v", p.name)
}

func (p person) clone() person {
	return person{p.name, p.age}
}

/* func (p person) someFunc() human{
	if p.name != "doesnotexist"{
		return nil
	}
	fmt.Println("creating a clone and returning...")
	return p.clone()
} */

var (
	p human = &person{}
)

func main() {
	// try 1
	/* p.name = "foo"
	p.age = 23
	p.speak() */
	// ./prog.go:23:3: p.name undefined (type human has no field or method name)
	// ./prog.go:24:3: p.age undefined (type human has no field or method age)

	// try 2 :  works
	p = &person{"bar", 23}
	fmt.Println(p.speak())

	// try 3 :  works
	clonedPerson := p.clone()
	fmt.Println(clonedPerson.speak())

	// try 4 :
	/* anotherClone := p.someFunc()
	if anotherClone != nil {
		fmt.Println(anotherClone.speak())
	} */
	// ./prog.go:52:19: p.someFunc undefined (type human has no field or method someFunc)

	// try 5
	newPerson := person{"johndoe", 34}
	if msg := newPerson.someFunc(); msg == nil {
		fmt.Println("able to return interface")
	} else {
		fmt.Println("not able to return interface")
	}
	// newPerson.someFunc undefined (type person has no field or method someFunc)
}
