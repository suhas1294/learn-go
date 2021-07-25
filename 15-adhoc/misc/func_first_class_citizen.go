package main

import(
  "fmt"
)

type User struct{
  Name string
  age int
}

func (u *User) GetUsername(s string) string {
  return (u.Name + s)
}

func funcReceiver(f func(s string) string) string {
  return f("SomeDefaultName")
}

func main() {

 u := User{
  "Foobar",
  25,
 }

 fmt.Println(funcReceiver(u.GetUsername))
 fmt.Println(u.GetUsername("Foobar"))

}