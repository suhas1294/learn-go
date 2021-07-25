package main

import(
	"fmt"
)

type Person struct{
	Id int32
	Name string
}


func method1(p *Person) {
	fmt.Println("---->",p.Name)
	p.Name = "universalName"
}

// mixed method
func method2(p Person) {
	method1(&p)
}

// mixed method
func method3(p Person, pp *Person) {
	// 'p' is a copy of original object
	fmt.Println("&p value:",&p, "pp value:", pp)
	fmt.Println("&p == pp ?",&p == pp)
	fmt.Println("address of object:")
	printAddressOfPerson(&p)
	fmt.Println("direct address:", pp)
	method1(&p)
}

func printAddressOfPerson(p *Person){
	fmt.Println(p)
}

func main(){
	//experiment 1 : directly sending person's address
	p2 := &Person{}
	method1(p2)
	fmt.Println("P2 name:",p2.Name)

	// experiment 2 : sending person object, whose address will be sent to method 1
	p4 := Person{
		78,
		"Foobar",
	}
	method2(p4)
	fmt.Println("p4 name",p4.Name)

	// p := Person{
	// 	5423,
	// 	"foobar",
	// }
	// method1(&p)
	// fmt.Println("P name:",p.Name)

	method3(p4, &p4)
	fmt.Println(p4.Name)

}