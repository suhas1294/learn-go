1. Common string methods
```
newString := strings.replace(emailBody, "${name}", username, -1) // replace all

```

2. creating a class with embedded type
```
type person struct {
	name string
	age  int
}

type employee struct {
	person
	salary float64
	skills []string
	isMarried bool
}
```
3. creating instance of a class
```
p := person{
	name: "foo",
	age:  24,
}

e := employee{
	person: person{
		"bar",
		35,
	},
	salary: 1200.50,
	skills: []string{"golang", "java", "c"},
	isMarried: true,
}

// accessing the fields of a type(accessing fields of an instance)
fmt.Println(e.name)
fmt.Println(e.skills)

```

4. There are two ways in which function returns : Pointers v/s Values.
   These are generic functions which can be compared to static methods in java.
   On the other hand, class-methods/instance methods are called `methods` in go.

    useful article : https://stackoverflow.com/questions/32208363/returning-value-vs-pointer-in-go-constructor
```
func createDefaultPerson() person{
	return person{
		"foo",
		23,
	}
}

func createDefaultPerson() *person{
	return &person{
		"foo",
		26,
	}
}
```
5. Internal conversion of address to value (i.e., pointer to type -> value of type) internal conversion :

    useful link: https://stackoverflow.com/questions/66796321/why-dereferencing-on-a-address-gives-invalid-indirect-error-in-golang

    Similar behaviour with maps: https://play.golang.org/p/cw3fq33qTSt
```
func printName(p *person){
    fmt.Println(p.name)
    // internally this will be understood as: fmt.Println( (*p).name )
}
p = person{"foo", 26}
printName(&p)
   ```
6. Class methods / instance methods :

    You have to attach methods to either `pointer receivers` or `value receivers`.

    useful articles :
    1. https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values
    2. https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver
```
// way 1 : pointer receiver
func (p *person) speak() {
	fmt.Println("My name is", p.name)
}
p = person{"foo", 26}
p.speak()

// way-2 : non-pointer receiver / value receiver
// here, func will be executed on copy of object on which method is called
func (p person) speak() {
	fmt.Println("My name is", p.name)
}
p = person{"foo", 26}
p.speak()

```

7. similarly, we can also return either values or pointers from a method:
```
func (p *person) constructEmployee() *employee{
	return &employee{
		person: person{
			p.name,
			p.age,
		},
		salary: 25000,
		skills: []string{"java", "python"},
		isMarried: true,
	}
}

human := createDefaultPerson()
emp := human.constructEmployee()
fmt.Println((*emp).name)

// ________________________________________

func (p *person) constructEmployee() employee{
	return employee{
		person: person{
			p.name,
			p.age,
		},
		salary: 25000,
		skills: []string{"java", "python"},
		isMarried: false,
	}
}
human := createDefaultPerson()
emp := human.constructEmployee()
fmt.Println(emp)
```