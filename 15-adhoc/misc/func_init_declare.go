package main

import "fmt"

type Person struct{
	Name string
	age int
}

var(
	foo func(p *Person) string
)

func main(){
	p := &Person{"Bar", 34}
	// if below function is uncommented, then foo(p) wont run
	// we are overriding a variable that was declared as a "function without logic" in var block.
	foo = func(p *Person) string{
		fmt.Println("Inside overrided func, value of p's name is", p.Name)
		return p.Name
	}
	foo(p)
	fmt.Println("about to end")
}
