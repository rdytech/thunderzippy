package main

import (
  "log"
  "net/http"
)

func main() {
  log.Printf("Thunderzippy is go on port %s", config.Port)
  http.HandleFunc("/zip/", ZipReferenceHandler)
  http.ListenAndServe(":" + config.Port, nil)
}

func HandleError(err error) {
  if err != nil {
    panic(err)
  }
}