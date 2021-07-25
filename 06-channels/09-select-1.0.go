package main

import (
  "fmt"
)

func main() {
  even := make(chan int)
  odd := make(chan int)
  quit := make(chan int)

  go send(even, odd, quit)

  receive(even, odd, quit)

  fmt.Println("about to exit")
}

// send channel
func send(even, odd, quit chan<- int) {
  for i := 0; i < 100; i++ {
    if i%2 == 0 {
      even <- i
    } else {
      odd <- i
    }
  }
  close(even)
  close(odd)
  quit <- 0
}

// receive channel
func receive(even, odd, quit <-chan int) {
  for {
    select {
    case v := <-even:
      fmt.Println("the value received from the even channel:", v)
    case v := <-odd:
      fmt.Println("the value received from the odd channel:", v)
    case v := <-quit:
      fmt.Println("the value received from the quit channel:", v)
      return
    }
  }
}
