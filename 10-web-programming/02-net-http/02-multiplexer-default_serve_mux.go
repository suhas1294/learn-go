/*
background:

1. Here in this example, if we pass'nil' has handler in ListenAndServe()
it will use DefaultServeNux().
2. So intead of ourMux.Handle(), we use http.Handle()

*/

package main

import(
  "io"
  "net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "dog dog dog !")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "cat cat cat !")
}

func main() {
  var d hotdog
  var c hotcat

  http.Handle("/dog", d)
  http.Handle("/cat", c)

  http.ListenAndServe(":8080", nil)

}