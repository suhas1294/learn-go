package main

import (
  "fmt"
)

func main() {
  c := make(chan int)

  go send(c)

  receive(c)

  close(c) // WE CLOSED THE CHANNEL

  fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
  c <- 42
}

// receive channel
func receive(c <-chan int) {
  fmt.Println("the value received from the channel:", <-c)
}
