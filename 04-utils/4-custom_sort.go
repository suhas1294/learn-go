package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// to-string implementation in go lang
func (personObj Person) toString() string {
	return fmt.Sprintf("%s: %d", personObj.Name, personObj.Age)
}

type SortPersonByAge []Person
type SortPersonByName []Person

// ______________________________
// Need to implement these methods: Len(), Swap(), Less()
// If a (any/custom) type has these three methods, it implicitly implements the (builtIn)'Interface' from package sort

/*
receiver 		: SortPersonByAge
function name 	: Len()
return type 	: int
*/
func (obj SortPersonByAge) Len() int {
	return len(obj)
}

/*
receiver 		: SortPersonByAge
function name 	: Swap()
return type 	: void
*/
func (obj SortPersonByAge) Swap(i, j int) {
	obj[i], obj[j] = obj[j], obj[i]
}

/*
receiver 		: SortPersonByAge
function name 	: Less()
return type 	: bool
*/
func (obj SortPersonByAge) Less(i, j int) bool {
	return obj[i].Age < obj[j].Age
}

// sorting by name
func (obj SortPersonByName) Len() int           { return len(obj) }
func (obj SortPersonByName) Swap(i, j int)      { obj[i], obj[j] = obj[j], obj[i] }
func (obj SortPersonByName) Less(i, j int) bool { return obj[i].Name < obj[j].Name }

// ______________________________

func main() {
	person1 := Person{"John Doe", 38}
	person2 := Person{"George bill", 24}
	person3 := Person{"Mathew", 54}
	person4 := Person{"Zenkusky", 16}

	people := []Person{person1, person2, person3, person4}

	fmt.Println(people) // prints the complete list

	for _, singlePerson := range people { // prints the people in list one by one
		fmt.Println(singlePerson)
	}

	// custom sorting - By Age
	sort.Sort(SortPersonByAge(people))
	fmt.Println("After sorting by age:\t", people)

	// custom sorting - By name
	sort.Sort(SortPersonByName(people))
	fmt.Println("After sorting by name:\t", people)
}
