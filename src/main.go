package main

import (
    "net/http"
    "bytes"
    "sync"
)

func makeHandler() func(w http.ResponseWriter, r *http.Request) {
  var b bytes.Buffer
  var mutex = &sync.Mutex{}

  return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
        b.WriteTo(w);
      case "PUT":
      case "POST":
        mutex.Lock()
        defer mutex.Unlock()
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
