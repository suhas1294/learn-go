/*
background:

1. type Handler interface has ServeHTTP(ResponseWriter, *Request) method.
2. Anything which implements ServeHTTP method is of type handler
3. In this example, "hotdog", "hotcat" -> both are of type hotdog and hotcat respectively and
ther underlying type is int. and they implement ServeHTTP, So they are of type Handle too.
4. Since "hotdog" and "hotcat" both are of type Handler, we can pass it to ListenAndServe method which requires a handler
5. Next Level - Multiplexer: There is a method NewServerMux() which will return the a (*ServeMux)
6. (*ServeMux has got so many methods attached to it, Handle() is one of them. )
7. Additionally (*ServeMux) implements Handler interface, hence, Any Object of type *ServeMux is also of type Handler and
hence can be passed to ListenAndServe()

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

  mux := http.NewServerMux()
  mux.Handle("/dog", d)
  mux.Handle("/cat", c)

  http.ListenAndServe(":8080", mux)

}