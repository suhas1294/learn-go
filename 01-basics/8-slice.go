package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}

	for index, value := range x {
		fmt.Println(index, value)
	}

	// slicing a slice
	// goes upto BUT NOT including
	y := x[1:3]
	fmt.Println(y)

	// Adding (APENDING) elements to a slice
	z := append(x, 12, 13, 99)
	fmt.Println(z)

	alpha := append(y, z...)
	fmt.Println(alpha)

	// DELETING from a slice
	delete := append(alpha[:2], alpha[4:]...)
	fmt.Println(delete)

	// use 'make' when u already know approximate capicity of slice
	cc := make([]int, 10, 100) // 10 is initial number of elements, 100 is capcity of slice
	fmt.Println(cc)
	fmt.Println(len(cc))
	fmt.Println(cap(cc))

	// cc[10] = 999 // error ! index out of range
	cc = append(cc, 999)
	fmt.Println(cc)

	// Multi dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
