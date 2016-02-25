package main

import (
  "os"
  "encoding/json"
  "log"
)

var config = Configuration{}

func init() {
  conf := "thunderzippy_conf.json"
  configFile, _ := os.Open(conf)
  log.Printf("Reading config %s", conf)
  decoder := json.NewDecoder(configFile)
  err := decoder.Decode(&config)
  if err != nil {
    panic("Error reading conf")
  }
}
