package main

import "fmt"

type human interface{
	Ping() string
}

type male struct{}

func (m *male) Ping() string{
	return "pong"
}

// create a variable of type interface
var foo human

// create a variable of type struct which implements interface, this works.
// var bar male

func main(){
	// calling method on interface, does not work
	fmt.Println(foo.Ping())
	
	// this works
	// fmt.Println(bar.Ping())
}
