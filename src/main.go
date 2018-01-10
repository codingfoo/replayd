package main

import (
    "net/http"
    "bytes"
)

func makeHandler() func(w http.ResponseWriter, r *http.Request) {
  var b bytes.Buffer

  return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
        b.WriteTo(w);
      case "PUT":
      case "POST":
        b.Reset()
        b.ReadFrom(r.Body);
      default:

    }
  }

}

func main() {
  handler := makeHandler()
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
