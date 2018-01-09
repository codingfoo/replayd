package main

import (
    "strings"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    datastr := "Some Data"
    data := strings.NewReader(datastr)
    postreq, err := http.NewRequest("POST", "/", data)
    if err != nil {
        t.Fatal(err)
    }

    postrr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler)

    handler.ServeHTTP(postrr, postreq)

    if status := postrr.Code; status != http.StatusOK {
        t.Errorf("handler returned the wrong status code for post request: got %v expected %v",
            status, http.StatusOK)
    }

    getreq, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    getrr := httptest.NewRecorder()

    handler.ServeHTTP(getrr, getreq)

    if status := getrr.Code; status != http.StatusOK {
        t.Errorf("handler returned the wrong status code for get request: got %v expected %v",
            status, http.StatusOK)
    }

    bodybytes, err2 := ioutil.ReadAll(getrr.Body)
    if err2 != nil {
        t.Fatal(err2)
    }

    body := string(bodybytes)

    if datastr != body {
      t.Errorf("Data did not match: got %v expected %v", body, datastr)
    }
}
