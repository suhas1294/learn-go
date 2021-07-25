package main

import (
  "log"
  "os"
  "text/template"
)

func main() {
  pointer_to_templates, err := template.ParseFiles("./../html-pages/01-html_for_parsing.gohtml", "file2.html",)
  if err != nil {
    log.Fatalln(err)
  }

  err = pointer_to_templates.ExecuteTemplate(os.Stdout, "file2.hmtl", nil) // observe the change here: ExecuteTemplate
  if err != nil {
    log.Fatalln(err)
  }

}

// pointer_to_templates, err := template.parseGlob("./../html-pages/*.html")