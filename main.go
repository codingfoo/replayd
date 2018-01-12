package main

import (
  "io/ioutil"
  "encoding/json"
  "flag"
  "log"
	"bytes"
	"net/http"
	"sync"
)

func makeHandler() func(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	var mutex = &sync.Mutex{}

	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		switch r.Method {
		case "GET":
			bytes.NewBuffer(b.Bytes()).WriteTo(w)
		case "PUT":
			fallthrough
		case "POST":
			b.Reset()
			b.ReadFrom(r.Body)
		default:
		}
	}
}

func main() {
  configFilePtr := flag.String("config-file", "config.json", "Path to the config file for the replayd daemon")
  flag.Parse()
  log.Printf("Config file path: %s", *configFilePtr)

  dat, err := ioutil.ReadFile(*configFilePtr)
  if err != nil {
    log.Printf("Error opening config file.")
    panic(err)
  }

  var config map[string]interface{}

  if err = json.Unmarshal(dat, &config); err != nil {
    panic(err)
  }

  host := config["host"].(string)
  port := config["port"].(string)

  log.Printf("Host: %s", host)
  log.Printf("Port: %s", port)

	handler := makeHandler()
	http.HandleFunc("/", handler)
  err = http.ListenAndServe(host + ":" + port, nil)
  if err != nil {
    log.Printf("ListenAndServer error: %s", err);
  }

}

