package main

import (
    "fmt"
    "net/http"
)

func hi_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[4:])
}

func main() {
    http.HandleFunc("/hi/", hi_handler)
    http.ListenAndServe(":8080", nil)
}
