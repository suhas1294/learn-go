package main

import (
  "encoding/json"
  "fmt"
)

type A struct {
  Foo string
  Bar *B
}

type B struct {
  Baz string
}

func main() {
  serialized := `{"foo": "some", "bar": {"baz": "thing"}}`
  a := &A{}
  json.Unmarshal([]byte(serialized), a)
  fmt.Println(a.Foo, a.Bar.Baz)

  // ________________________________

  myStructExample := MyStruct{
    Val: 23,
  }
  fmt.Println(myStructExample)
  fmt.Println(&myStructExample)
  fmt.Printf("type is : %T\n\n", &myStructExample)

  fmt.Println("creating a dummy object without any values", MyStruct{}) // same this is being done in myfunc()

  myStructObj := myfunc()     // got the address of (new MyStruct object with its field having zero values)
  *myStructObj = MyStruct{43} // changing the value stored in that address
  // *myStructObj = B{"not possible"}
  fmt.Println(*myStructObj) // printing the value stored in that address.
}

// ________________________________

type MyStruct struct {
  Val int
}

func myfunc() *MyStruct {
  return &MyStruct{} // constructing anonymous object of type MyStruct and returning it, field will have zero values
}
