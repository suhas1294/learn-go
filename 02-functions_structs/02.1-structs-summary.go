// (instance of class) == (value of type)
type person struct{
	first string
	last string
}

p1 := person{
	first: "james",
	last: "bond"
}

// ______________________________________________________
// embedded struct

type secretAgent struct{
	person // anonymous field
	ltk bool
}

sa1 := secretAgent{
	person: person{
		first: "James",
		last:  "Bond",
		age:   32,
	},
	ltk: true,
}

fmt.Println(sa1)
fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

// ______________________________________________________
// anonymous struct

p1 := struct{
	first string
	last string
	age int	
}{
	first: "james",
	last: "bond",
	age: 32
}