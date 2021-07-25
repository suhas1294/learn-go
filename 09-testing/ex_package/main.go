package main

import (
  "fmt"
  "./acdc"
)

func main() {
  fmt.Println(acdc.Sum(2, 3))
  fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}