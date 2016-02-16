package main

import (
    "io"
    "net/http"
    "log"
    "archive/zip"
    "time"
)

type ZipEntry struct {
    Filepath, Url string
}

func zip_entries() *[]ZipEntry{
    list := make([]ZipEntry, 2)
    list[0] = ZipEntry{
        "images/CC-attribution.png",
        "http://localhost:3000/images/CC-attribution.png",
    }
    list[1] = ZipEntry{
        "images/facebook-small.png",
        "http://localhost:3000/images/facebook-small.png",
    }
    return &list
}

func zip_handler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    w.Header().Add("Content-Disposition", "attachment; filename=\"test.zip\"")
    w.Header().Add("Content-Type", "application/zip")

    zipWriter := zip.NewWriter(w)

    for _, zip_entry := range *zip_entries() {
        log.Printf("Get:\t%s", zip_entry.Url)
        add_download_to_zip(zipWriter, zip_entry.Url, zip_entry.Filepath)
    }

    err := zipWriter.Close()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start))
}

func add_download_to_zip(zipWriter *zip.Writer, url string, file_path_inside_zip string) {
    resp, err := http.Get(url)
    defer resp.Body.Close()
    if err != nil {
        log.Fatal(err)
        return
    }

    // https://golang.org/pkg/archive/zip/#FileHeader
    h := &zip.FileHeader{
        Name:   file_path_inside_zip,
        Method: zip.Deflate,
        Flags:  0x800,
    }
    h.SetModTime(time.Now())
    f, err := zipWriter.CreateHeader(h)
    if err != nil {
        log.Fatal(err)
    }

    io.Copy(f, resp.Body)
}

func main() {
    log.Printf("Thunderzippy is go")
    http.HandleFunc("/zip/", zip_handler)
    http.ListenAndServe(":8080", nil)
}
