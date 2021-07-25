package main

import(
  "fmt"
  "log"
  "os"
)

func main() {
  f, err := os.Create("log.txt")
  if (err != nil) {
    fmt.Println("Error creating a log file which where errors are logged.")
  }
  defer f.Close()
  log.SetOutput(f) // we are telling logger to log to above created file

  f2, err := os.Open("non_existing_file.txt")
  if (err != nil)  {
    log.Println("Err happened :", err)
  }
  defer f2.Close()
  fmt.Println("Check the log.txt file for error tracing")
}