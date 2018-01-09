package main

import (
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusTeapot)
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
