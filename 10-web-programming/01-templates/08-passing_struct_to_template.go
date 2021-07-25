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

type Sage struct{
  Name string
  Motto string
}

type Car struct{
  Name string
  Model string
  Doors int
}

type Items struct{
  Wisdom []Sage
  Transport []Car
}

func main() {

  struct_data := Sage{
    Name: "Vivekananda",
    Motto: "Awake, arise until you succeed !",
  }


  // For data of type struct
  err := pointer_to_templates.ExecuteTemplate(os.Stdout, "template_for_struct.gohtml", struct_data)
  if err != nil {
    log.Fatalln(err)
  }

  // For data of type 'slice of struct'

  sage1 := Sage{"Vishwamitra", "Gayatri Mantra"}
  sage2 := Sage{"Agastya", "Kriya yoga"}
  sage3 := Sage{"Osho", "Jnana yoga"}

  slice_of_sages := []Sage{sage1, sage2, sage3}

  err = pointer_to_templates.ExecuteTemplate(os.Stdout, "template_for_sliceStruct.gohtml", slice_of_sages)
  if err != nil {
    log.Fatalln(err)
  }

  // For data of type 'struct containing struct'
  car1 := Car{"Ferrari", "GTC4", 2,}
  car2 := Car{"Tata", "Indica", 4,}
  car3 := Car{"Tesla", "model S", 5,}

  slice_of_cars := []Car{car1, car2, car3}

  struct_of_struct := Items{
    Wisdom: slice_of_sages,
    Transport: slice_of_cars,
  }

  // this can also be done via anonymous struct
  data_set2 := struct{
    Wisdom []Sage
    Transport []Car
  }{
    slice_of_sages,
    slice_of_cars,
  }

  //err = pointer_to_templates.ExecuteTemplate(os.Stdout, "struct_of_struct_template.gohtml", struct_of_struct)
  err = pointer_to_templates.ExecuteTemplate(os.Stdout, "struct_of_struct_template.gohtml", data_set2)
  if err != nil {
    log.Fatalln(err)
  }

}