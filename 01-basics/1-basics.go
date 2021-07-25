// https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#
// https://www.youtube.com/watch?v=XCsL89YtqCs
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}

// variadic parameters
// take a value 'a' of any type (interface{}) and unlimited number of them (...)
func Println(a ...interface{}) (n int, err error)