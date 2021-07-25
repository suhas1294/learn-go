package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

	// -----------------------------
	// Func  expressions, func's are first class citizens, they can be assigned to variable
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)
	// -----------------------------
}

func foo() {
	fmt.Println("foo ran")
}
