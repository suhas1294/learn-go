package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  // value of a
	fmt.Println(&a) // address of a

	fmt.Printf("%T\n", a)  // type is int
	fmt.Printf("%T\n", &a) // type is *int (pointer to int)

	//var b int = &a // Will throw error (bcoz, &a type is pointer to int, cannot be assigned to an int type variable)
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // b's value is 42's address, so (pointer to address gives values)

	/*
		pointer to type vs pointer to address
		*int : asterisk with type : is pointer to the type
		*<some_variable> : asterisk with variable : is pointer to the type address - dereferencing
		deferencing : when a address is stored in a variable, if we do *<variable_that_has_adress>,
		then it will give value stored at that address.
		'&' gives the address
		'*' gives the value stores in address

		Usage:
			when u have a large chunk of data, instead of passing it accross, u can pass its address
	*/
}
