package main

import(
  "fmt"
)

func main() {
  example_channel := make(chan int)
  go func(){
      example_channel <- 42
    }()
  fmt.Println(<-example_channel)
}