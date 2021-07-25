package main

import (
  "fmt"
)

type person struct {
  first string
  last  string
  age   int
}

type Human struct{
  person
  Senior bool
}

func main() {
  human := Human{
    person: person{
      first: "Foo",
      last: "Bar",
      age: 25,
    },
    Senior: true,
  }

  fmt.Println(human.first, human.last, human.age, human.Senior)
  // even though 'first', 'last' and 'age' are not direct fields of tpye 'human', when called it gets internally promoted to outer type
  // So we will be able to access the fields of inner types via outer type
  // this is called as Inner promotion
}
