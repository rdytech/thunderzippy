package main

type Configuration struct {
  RedisServerAndPort string
  Port               string
}

type ZipEntry struct {
  Filepath, Url string
}