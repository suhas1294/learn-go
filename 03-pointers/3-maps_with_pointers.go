package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

var (
	pm1 = map[string]*Person{
		"p1": {Name: "Foo", Age: 25},
	}
	// below commented code does not work !
	/* pm2 = map[string]*Person{
		"p2": &{Name: "bar", Age: 23},
	} */
)

func main() {
	fmt.Println(pm1)
	// fmt.Println(pm2) // does not work

	p := pm1["p1"]
	fmt.Printf("type of p is %T\n", p)
	fmt.Println(p.Name)
	fmt.Println((*p).Name)
}
