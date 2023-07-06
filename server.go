package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    originalQuery := ""

    if queryBy {
      filters := ""
    } else {
      filters := ""
    }

    query := fmt.Sprintf("%s %s", originalQuery, filters)
    fmt.Fprintf(w, returnFields)
}
