package main

import(
  "fmt"
)

// concurrency is a design pattern
// parallelism is a actual execution
// share memory by communicating

func main() {
  example_channel := make(chan int)
  example_channel <- 42
  fmt.Println(<-example_channel)
}