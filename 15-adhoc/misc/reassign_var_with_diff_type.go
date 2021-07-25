package main

import "fmt"

type Human interface{
	Speak(string) string
}

type Person struct{
	Name string
	age int
}

func (p *Person) Speak(s string) string{
	return fmt.Sprintf("Speaking %v", s)
}

var(
	h Human
	p Person
)

func main(){
	// observe that h was initialised as a variable of type Human interface.
	// But we are able to assign a 'pointer to person' type to h variable.
	// this is not possible if we had initialized h as Person type instead of interface type.
	// uncomment below code to understand how it breaks the flow.
	h = &Person{
		Name: "foo",
		age: 23,
	}
	fmt.Printf("type of h is %T\n", h)
	fmt.Println("value of h", h)

	// below code does not work
	/* p = &Person{"bar", 36}
	fmt.Printf("type of p is %T\n", p)
	fmt.Println(p) */
	
}
