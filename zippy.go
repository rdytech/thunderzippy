package main

import (
    "archive/zip"
    "io"
    "log"
    "net/http"
    "time"
)

type ZipEntry struct {
    Filepath, Url string
}

func zip_entries() *[]ZipEntry {
    list := make([]ZipEntry, 2)
    list[0] = ZipEntry{
        "images/CC-attribution.png",
        "http://localhost:3000/images/CC-attribution-not-found.png",
    }
    list[1] = ZipEntry{
        "images/facebook-small.png",
        "http://localhost:3000/images/facebook-small.png",
    }
    return &list
}

func zip_handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s\t\t%s", r.Method, r.RequestURI)
    start := time.Now()
    w.Header().Add("Content-Disposition", "attachment; filename=\"test.zip\"")
    w.Header().Add("Content-Type", "application/zip")

    zipWriter := zip.NewWriter(w)

    for _, zip_entry := range *zip_entries() {
        add_download_to_zip(zipWriter, zip_entry.Url, zip_entry.Filepath)
    }

    err := zipWriter.Close()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Tzipped:\t(%s)", time.Since(start))
}

func add_download_to_zip(zipWriter *zip.Writer, url string, name string) {
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

func main() {
    log.Printf("Thunderzippy is go")
    http.HandleFunc("/zip/", zip_handler)
    http.ListenAndServe(":8080", nil)
}
