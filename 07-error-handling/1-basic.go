/*
type error interface{
  Error() string
}
*/

// any type that has Error() is also an error.
// just like 'int', 'float64', 'error' is also an type

// with panic, defer function runs, recover can be used,
// But with 'fatal' , system just shuts down

package main

import (
  "log"
  "os"
)

func main() {
  _, err := os.Open("no-file.txt")
  if err != nil {
    //    fmt.Println("err happened", err)
    log.Println("err happened", err)
    //    log.Fatalln(err)
    //    panic(err)
  }
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ...
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.