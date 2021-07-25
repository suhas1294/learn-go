package main

import (
  "fmt"
)

// can be used to know if a channel is closed
// usually used while receiving from channel
func main() {
  c := make(chan int)
  go func() {
    c <- 42
  }()

  v, ok := <-c

  fmt.Println(v, ok)
}
