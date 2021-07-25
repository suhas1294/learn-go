package main

import (
  "fmt"
)

func main() {
  c := make(chan int)

  go send(c)

  receive(c)

  fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
  for i := 0; i < 100; i++ {
    c <- i
  }
  close(c) // if this is not closed, then deadlock problem may arise in the receiving end.
}

// receive channel
func receive(c <-chan int) {
  for v := range c { // ranging untill we close (c) in send()
    fmt.Println("the value received from the channel:", v)
  }
}
