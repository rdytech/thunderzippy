package main

import (
  "log"
  "net/http"
  "os"
)

func main() {
  log.Printf("Thunderzippy is go on port %s", os.Getenv("PORT"))
  http.HandleFunc("/zip/", ZipReferenceHandler)
  http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}

func HandleError(err error) {
  if err != nil {
    panic(err)
  }
}
