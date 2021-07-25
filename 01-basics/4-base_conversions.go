package main

import "fmt"

func main() {
	s := "Hello World"
	bs := []byte(s)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("UTF value is\t%#U\n", s[i])
	}
}
