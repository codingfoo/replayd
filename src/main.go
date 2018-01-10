package main

import (
	"bytes"
	"net/http"
	"sync"
)

func makeHandler() func(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	var mutex = &sync.Mutex{}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			bytes.NewBuffer(b.Bytes()).WriteTo(w)
		case "PUT":
			fallthrough
		case "POST":
			mutex.Lock()
			defer mutex.Unlock()
			b.Reset()
			b.ReadFrom(r.Body)
		default:
		}
	}
}

func main() {
	handler := makeHandler()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
