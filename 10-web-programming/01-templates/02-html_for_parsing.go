//Take away: Parsing and execution of templatespackage main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  pointer_to_template, err := template.ParseFiles("./../html-pages/01-html_for_parsing.gohtml")
  // pointer_to_template, err := template.ParseFiles("./../01-html_for_parsing.gohtml", "file2.html", ..) // can hold as many templates as u want
  // pointer_to_template is like a container which hold all the data that we have given for parsing
  if err != nil {
    log.Fatalln(err)
  }
  err = pointer_to_template.Execute(os.Stdout, nil) // Execute takes writer and data, in this case data is nil
  if err != nil {
    log.Fatalln(err)
  }

}


/*
// variant program (write output to file)

package main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  pointer_to_template, err := template.ParseFiles("./../01-html_for_parsing.gohtml")
  // pointer_to_template, err := template.ParseFiles("./../01-html_for_parsing.gohtml", "file2.html", ..) // can hold as many templates as u want
  // pointer_to_template is like a container which hold all the data that we have given for parsing
  if err != nil {
    log.Fatalln(err)
  }

  new_file, err := os.Create("index.html")
  if err != nil{
    log.Println("Not able to create file")
  }
  defer new_file.close

  err = pointer_to_template.Execute(new_file, nil) // observe the change here
  if err != nil {
    log.Fatalln(err)
  }

}
*/