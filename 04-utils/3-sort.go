package main

import (
	"fmt"
	"sort"
)

func main() {

	int_slice := []int{2, 334, 4, 123, 536, 434, 12, 5, 68, 3, 467, 3445, 567}
	string_slice := []string{"Sughosh", "Abhiram", "chintan", "karthik", "Uganda"}

	sort.Ints(int_slice)       // no need of reassigning back to slice
	sort.Strings(string_slice) // no need of reassigning back to slice

	fmt.Println(int_slice)
	fmt.Println(string_slice)
}
