package main

import(
  "fmt"
)

func main() {
  example_channel := make(chan int, 1) // my buffer channel has 1 to sit in there regardless of anyone pulling it off
  example_channel <- 42 // this is not happening in different thread, so no one os blocking it
  fmt.Println(<-example_channel)
}