package main

import (
  "encoding/json"
  "strconv"
  "log"
  "time"
  "errors"
  "math/rand"
  redigo "github.com/garyburd/redigo/redis"
)

var redisPool *redigo.Pool

func init() {
  log.Printf("Opening Redis: %s", config.RedisServerAndPort)
  redisPool = &redigo.Pool{
    MaxIdle:     10,
    IdleTimeout: 1 * time.Second,
    Dial: func() (redigo.Conn, error) {
      return redigo.Dial("tcp", config.RedisServerAndPort)
    },
    TestOnBorrow: func(c redigo.Conn, t time.Time) (err error) {
      _, err = c.Do("PING")
      if err != nil {
          panic("Error connecting to redis")
      }
      return
    },
  }
}

func getFileListByZipReferenceId(id string) (files []*ZipEntry, err error) {
  redis := redisPool.Get()
  defer redis.Close()
  
  // Get the value from Redis
  result, err := redis.Do("GET", "zip:" + id)
  if err != nil || result == nil {
    err = errors.New("Access Denied (sorry your link has timed out)")
    return
  }

  // Convert to bytes
  var resultByte []byte
  var ok bool
  if resultByte, ok = result.([]byte); !ok {
    err = errors.New("Error converting data stream to bytes")
    return
  }

  // Decode JSON
  if err = json.Unmarshal(resultByte, &files); err != nil {
    err = errors.New("Error decoding json: " + string(resultByte))
  }
  return
}

func CreateZipReference(files []*ZipEntry) (ref_id_string string) {
  redis := redisPool.Get()
  defer redis.Close()
  
  filesJson, err := json.Marshal(files)
  HandleError(err)
  
  //get new id redis
  ref_id, err := redis.Do("INCR", "zip_reference_id")
  HandleError(err)
  ref_id_string = RandomString(5) + strconv.FormatInt(ref_id.(int64), 10)

  // Save JSON files to Redis
  _, err = redis.Do("SET", "zip:" + ref_id_string, filesJson)
  HandleError(err)
  
  return
}

func RandomString(strlen int) string {
  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}