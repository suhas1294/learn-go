package main

import "fmt"

var a  int
type hotdog int // hotdog's underlying type is int.
var b hotdog

func main{
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a) //int
	
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b) // main.hotdog

	a = b // does not work

	a = int(b) // conversion
	fmt.Printf("%T\n", a) //int
}