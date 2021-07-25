package main

import "fmt"

type Person struct{
  Name string
  Age int
}

func (p *Person) pointerReceiverMethod() {
  p.Name = "p-Name"
  fmt.Println("This is called by ", p.Name)
}

func (p Person) valueReceiverMethod() {
  p.Name = "v-name"
  fmt.Println("Truth has been spoken by ", p.Name)
}

// Functions with return types

// changing the received person's value and returning a clone of updated person
func (p *Person) modifyPersonWay1() (*Person) {
  p.Name = "ret-p-Name-1"
  return &Person{p.Name, p.Age}
}

// changing the received person's value and returning a clone of updated person
func (p *Person) modifyPersonWay2() (*Person) {
  p.Name = "ret-p-Name-2"
  return p
}

func (p *Person) modifyPersonWay3() (Person) {
  p.Name = "ret-p-Name-3"
  return *p
}

func (p *Person) modifyPersonWay4() (Person) {
  p.Name = "Tesla"
  return *p
}

func (p *Person) test1(){
  fmt.Printf("%T\n", p)
  fmt.Printf("%T\n", &p)
  fmt.Printf("%T\n", *p)
}

func main() {
  person1 := Person{
    "Foo bar",
    25,
  }

  person1.valueReceiverMethod()
  fmt.Println("person1's name has changed ?", person1.Name)

  person1.pointerReceiverMethod()
  fmt.Println("person1's name has changed ?", person1.Name)

  person1.test1()

  fmt.Println("_______________________________________")

  person2 := Person{"Agastya", 108}
  fmt.Println("person2 Before:\t\t\t",person2.Name)
  modifiedPerson2 := person2.modifyPersonWay1()
  fmt.Println(modifiedPerson2)
  fmt.Println("person2 After:\t\t\t",person2.Name)
  fmt.Println("modifiedPerson2 (clone of person2):\t\t\t",modifiedPerson2.Name)

  fmt.Println("_______________________________________")

  person3 := Person{"Atreya", 109}
  fmt.Println("person3 Before:\t\t\t",person3.Name)
  modifiedPerson3 := person3.modifyPersonWay2()
  fmt.Println(modifiedPerson3)
  fmt.Println("person3 After:\t\t\t",person3.Name)
  fmt.Println("modifiedPerson3 (clone of person3):\t\t\t",modifiedPerson3.Name)

  fmt.Println("_______________________________________")

  person4 := Person{"Vishwa Mitra", 111}
  fmt.Println("person4 Before:\t\t\t",person4.Name)
  modifiedPerson4 := person4.modifyPersonWay3() // here person4's name changes first and then returned value of person4 is assigned back to new variable modifiedPerson4
  fmt.Printf("\nReturn type of modifiedPerson4 is %T\n", modifiedPerson4)
  fmt.Println("modifiedPerson4 initialised :\t\t\t",modifiedPerson4)
  fmt.Println("person4 After:\t\t\t",person4.Name)
  fmt.Println("modifiedPerson4 (clone of person4):\t\t\t",modifiedPerson4.Name)

  // Now if i change modifiedPerson4's name, it will reflect in person4
  modifiedPerson4.modifyPersonWay4()
  fmt.Println("person4 After updating modifiedPerson4:\t",person4.Name)


  fmt.Println("_______________________________________")

}