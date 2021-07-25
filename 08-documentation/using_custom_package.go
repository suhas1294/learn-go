package main

import (
  "fmt"
  "github.com/suhas1294/learn-go/8-documentation/customPackage"
  //"./customPackage" //while using local
)

func main() {
  fmt.Println("printing from custom package:")
  customPackage.CustomFunction()
}
