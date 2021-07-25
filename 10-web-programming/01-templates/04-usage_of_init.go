package main

import (
  "log"
  "os"
  "text/template"
)

var pointer_to_templates *template.Template

func init(){
  pointer_to_templates, err = template.parseGlob("./../html-pages/*.html")
  // pointer_to_templates = template.Must(template.parseGlob("./../html-pages/*.html")) // 'Must' does not return err
}

func main() {
  err := pointer_to_templates.ExecuteTemplate(os.Stdout, "file2.hmtl", nil) // observe the change here: ExecuteTemplate
  if err != nil {
    log.Fatalln(err)
  }

}