package main

import (
  "log"
  "os"
  "text/template"
)

var pointer_to_templates *template.Template

func init(){
  template_path := `/Users/username/workspace/backend/go/src/go_workspace/src/web-programming/01-templates/html-pages/*.gohtml`
  pointer_to_templates = template.Must(template.ParseGlob(template_path))
}

func main() {
  // err := pointer_to_templates.ExecuteTemplate(os.Stdout, "template_with_data.gohtml", 42) // passing the 'value of type int'
  err := pointer_to_templates.ExecuteTemplate(os.Stdout, "template_with_data.gohtml", "Some Data !!") // passing the 'value of type string'
  if err != nil {
    log.Fatalln(err)
  }

}