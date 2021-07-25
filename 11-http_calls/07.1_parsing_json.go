package main

import(
  //"encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
  "log"
  "github.com/tidwall/gjson"
)

const (
	PROTOCOL = "https"
	HOST = "app.swiggy.com"
	PATH = "/api/v4/restaurants/"
	PATH_PARAMS = "59802"
	LAT = "?lat=12.9569&"
	LNG = "?lng=77.7011"
)

func main() {
  menu_url := PROTOCOL + "://" + HOST + PATH + PATH_PARAMS + LAT + LNG
  resp, err := http.Get(menu_url)
  if err != nil {
    log.Fatalln(err)
  }
  defer resp.Body.Close()

  //fmt.Printf("\n%T\n", resp)


  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }

  //fmt.Println(string(body))
  //fmt.Printf("\n%T\n", body)

  /*
  var f interface{}
  err = json.Unmarshal(body, &f)
  fmt.Printf("\n%T\n\n", f)
  fmt.Println(f)
  */

  value := gjson.Get(string(body), "data.menu.items.15235618.ribbon.text")
  fmt.Println(value)

}