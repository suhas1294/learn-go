// what is callback : passing func as argument

package main

import (
	"fmt"
)

func main() {
	/*ii := []int{1,2,3,4,5,6,7,8,9,}
	s:= sum(ii...) // unfurling the slice
	fmt.Println("Sum of all numbers: ", s)*/
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func foo() func() int { // func 'foo' that returns an ('func' which will return an int)
	fmt.Println("")
}
