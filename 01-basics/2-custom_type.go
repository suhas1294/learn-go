package main

import (
	"fmt"
)

type hotdog int // int is the underlying type of new type hotdog

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // %T to get the type of variable
	x = 42
	fmt.Println(x)
}
