package main

import (
  "fmt"
)

func main() {
  c := make(chan int)
  cr := make(<-chan int) // receive
  cs := make(chan<- int) // send

  fmt.Println("-----")
  fmt.Printf("c\t%T\n", c)
  fmt.Printf("cr\t%T\n", cr)
  fmt.Printf("cs\t%T\n", cs)

  // specific to general doesn't assign - DOES NOT WORK !!
  c = cr
  c = cs

  // specific to specific doesn't assign - DOES NOT WORK !!
  cs = cr

  // general to specific assigns - WORKS !!
  cr = c
  cs = c
}
