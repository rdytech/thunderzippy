package main

import (
    "io"
    "net/http"
    "log"
    "archive/zip"
    "time"
)

func zip_handler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    w.Header().Add("Content-Disposition", "attachment; filename=\"test.zip\"")
    w.Header().Add("Content-Type", "application/zip")

    zipWriter := zip.NewWriter(w)

    url := "http://localhost:3000/images/CC-attribution.png"
    file_path_inside_zip := "images/CC-attribution.png"
    log.Printf("Get:\t%s", url)

    add_download_to_zip(zipWriter, url, file_path_inside_zip)
    zipWriter.Close()

    log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, time.Since(start))
}

func add_download_to_zip(zipWriter *zip.Writer, url string, file_path_inside_zip string) {
    h := &zip.FileHeader{
        Name:   file_path_inside_zip,
        Method: zip.Deflate,
        Flags:  0x800,
    }

    f, _ := zipWriter.CreateHeader(h)

    resp, _ := http.Get(url)
    defer resp.Body.Close()
    io.Copy(f, resp.Body)
}

func main() {
    http.HandleFunc("/zip/", zip_handler)
    http.ListenAndServe(":8080", nil)
}
