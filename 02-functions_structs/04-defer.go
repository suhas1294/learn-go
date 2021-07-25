package main

import (
	"fmt"
)

// defer = postpone
// classical example : write a method to close the file and define it in defer
// defer function will run after containing function is over

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
