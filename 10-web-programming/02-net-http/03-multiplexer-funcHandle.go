/*
background:

1. Here in this example, we use HandleFunc() which accepts url, func(ResponseWriter, *Request)
(Basically we are passing url and func as arguments to HandleFunc)

*/

package main

import(
  "io"
  "net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "dog dog dog !")
}

func c(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "cat cat cat !")
}

func main() {

  http.HandleFunc("/dog", d)
  http.HandleFunc("/cat", c)

  http.ListenAndServe(":8080", nil)

}