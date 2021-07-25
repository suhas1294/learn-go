package main

import(
  "fmt"
  "encoding/json"
)

type A struct {
    Foo string
    Bar B
}

type B struct {
    Baz string
}

func main() {
    serialized := `{"foo": "some", "bar": {"baz": "thing"}}`
    a := &A{}
    json.Unmarshal([]byte(serialized), a)
    fmt.Println(a.Foo, a.Bar.Baz)

    /*testObj := &B{
      Baz: "qwertyu",
    }
    fmt.Println(*testObj)

    testObj.Baz = "asdfghj"
    fmt.Println(*testObj)*/

}