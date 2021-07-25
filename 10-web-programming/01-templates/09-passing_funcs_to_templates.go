
/**

Supporting links:
  https://stackoverflow.com/questions/25169094/in-go-templates-i-can-get-parse-to-work-but-cannot-get-parsefiles-to-work-in-li
  https://golang.org/pkg/html/template/#New
  https://stackoverflow.com/questions/41176355/go-template-name
  https://stackoverflow.com/questions/49043292/error-template-is-an-incomplete-or-empty-template
  Error code : https://play.golang.org/p/FjI_jZOGQFB
*/
package main

import (
  "log"
  "os"
  "strings"
  "text/template"
  "path"
)

var tpl *template.Template

type sage struct {
  Name  string
  Motto string
}

type car struct {
  Manufacturer string
  Model        string
  Doors        int
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}


var filePath string
var fileErr error

func init() {
  filePath, fileErr = os.Getwd()
  if fileErr != nil {
       log.Fatal(fileErr)
  }
  filePath = filePath + "/html-pages/templates_with_funcs.gohtml"
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles(filePath))
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
  if len(s) >= 3 {
    s = s[:3]
  }
  return s
}

func (s sage) FuncAttachedToSage() string{
  return s.Name + ": called from type's method"
}

func main() {

  b := sage{
    Name:  "Buddha",
    Motto: "The belief of no beliefs",
  }

  g := sage{
    Name:  "Gandhi",
    Motto: "Be the change",
  }

  m := sage{
    Name:  "Martin Luther King",
    Motto: "Hatred never ceases with hatred but with love alone is healed.",
  }

  f := car{
    Manufacturer: "Ford",
    Model:        "F150",
    Doors:        2,
  }

  c := car{
    Manufacturer: "Toyota",
    Model:        "Corolla",
    Doors:        4,
  }

  sages := []sage{b, g, m}
  cars := []car{f, c}

  data := struct {
    Wisdom    []sage
    Transport []car
  }{
    sages,
    cars,
  }

  err := tpl.ExecuteTemplate(os.Stdout, path.Base(filePath), data) // You should always pass absolute fileName here not the filePath

  /*
  err := tpl.ExecuteTemplate(os.Stdout, path.Base(filePath), data)
  // Error : template: no template "/Users/username/workspace/backend/go/src/go_workspace/src/web-programming/01-templates/html-pages/templates_with_funcs.gohtml" associated with template ""
  */

  if err != nil {
    log.Fatalln(err)
  }
}