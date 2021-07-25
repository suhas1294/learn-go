package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 43
	fmt.Println(x)
	fmt.Println(len(x))
}
