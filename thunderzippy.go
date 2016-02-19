package main

import (
    "archive/zip"
    "encoding/json"
    "errors"
    redigo "github.com/garyburd/redigo/redis"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

var redisPool *redigo.Pool
var config = Configuration{}

type Configuration struct {
    RedisServerAndPort string
    Port               string
}

type ZipEntry struct {
    Filepath, Url string
}

func zipHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s\t\t%s", r.Method, r.RequestURI)
    start := time.Now()

    refs, ok := r.URL.Query()["ref"]
    if !ok || len(refs) < 1 {
        http.Error(w, "Thunderzippy. Pass ?ref= to use.", 500)
        return
    }
    ref := refs[0]

    files, err := getFileListFromRedis(ref)
    if err != nil {
        http.Error(w, err.Error(), 403)
        log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
        return
    }

    w.Header().Add("Content-Disposition", "attachment; filename=\"download.zip\"")
    w.Header().Add("Content-Type", "application/zip")

    zipWriter := zip.NewWriter(w)

    for _, file := range files {
        addDownloadToZip(zipWriter, file.Url, file.Filepath)
    }

    err = zipWriter.Close()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Thunderzipped:\t%d files (%s)", len(files), time.Since(start))
}

func addDownloadToZip(zipWriter *zip.Writer, url string, name string) {
    // https://golang.org/pkg/net/http/
    resp, err := http.Get(url)
    if err != nil {
        log.Print(err)
        return
    }
    log.Printf("adding:\t%d %s", resp.StatusCode, url)
    if resp.StatusCode != 200 {
        return
    }
    defer resp.Body.Close()

    // https://golang.org/pkg/archive/zip/#FileHeader
    h := &zip.FileHeader{
        Name:   name,
        Method: zip.Deflate,
        Flags:  0x800,
    }
    h.SetModTime(time.Now())
    f, err := zipWriter.CreateHeader(h)
    if err != nil {
        log.Print(err)
        return
    }

    io.Copy(f, resp.Body)
}

func getFileListFromRedis(ref string) (files []*ZipEntry, err error) {
    redis := redisPool.Get()
    defer redis.Close()

    // Get the value from Redis
    result, err := redis.Do("GET", "zip:"+ref)
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
    err = json.Unmarshal(resultByte, &files)
    if err != nil {
        err = errors.New("Error decoding json: " + string(resultByte))
    }

    return
}

func initRedis() {
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

func initConfig() {
    conf := "thunderzippy_conf.json"
    configFile, _ := os.Open(conf)
    log.Printf("Reading config %s", conf)
    decoder := json.NewDecoder(configFile)
    err := decoder.Decode(&config)
    if err != nil {
        panic("Error reading conf")
    }
}

func main() {
    initConfig()
    initRedis()
    log.Printf("Thunderzippy is go on port %s", config.Port)
    http.HandleFunc("/zip/", zipHandler)
    http.ListenAndServe(":"+config.Port, nil)
}
