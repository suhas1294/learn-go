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
  map_data := map[string]string{
    "India": "Incredible",
    "Bhutan": "Green",
    "Nepal": "Mountains",
  }

  err := pointer_to_templates.ExecuteTemplate(os.Stdout, "template_for_map.gohtml", map_data)
  if err != nil {
    log.Fatalln(err)
  }

}