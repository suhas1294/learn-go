package main

import "fmt"

// Interface says : Anyone who uses this method, is of My type !
// a value can be of more then one type
// ------------------------------------------------------
type human interface {
	speak()
}

// ------------------------------------------------------
// Models / inheriters
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// ------------------------------------------------------
// BOTH secretAgent and person has speak() method, SO they are are of HUMAN type
// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// ------------------------------------------------------
// utility function

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}
