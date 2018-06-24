package main

import (
  "archive/zip"
  "encoding/json"
  "net/http"
  "io"
  "io/ioutil"
  "time"
  "log"
)

func ZipReferenceHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      ZipReferenceDownload(w, r)
    case "POST":
      ZipReferenceCreate(w, r)
  }
}

func ZipReferenceDownload(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s\t\t%s", r.Method, r.RequestURI)
    start := time.Now()

    refs, ok := r.URL.Query()["ref"]
    if !ok || len(refs) < 1 {
      http.Error(w, "Thunderzippy. Pass ?ref= to use.", 500)
      return
    }
    ref := refs[0]

    files, err := getFileListByZipReferenceId(ref)
    if err != nil {
      http.Error(w, err.Error(), 403)
      log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, err.Error())
      return
    }

    w.Header().Add("Content-Disposition", "attachment; filename=\"documents-" + ref + ".zip\"")
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

func ZipReferenceCreate(w http.ResponseWriter, r *http.Request) {
  var files []*ZipEntry

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    http.Error(w, "Data was too limited to allow", 403)
    return
  }

  if err := r.Body.Close(); err != nil {
    http.Error(w, "Wrong Data. Please try again", 403)
    return
  }

  if err := json.Unmarshal(body, &files); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)

    if err := json.NewEncoder(w).Encode(err); err != nil {
      http.Error(w, "Wrong Data. Please try again", 403)
      return
    }
  }

  ref_id := CreateZipReference(files)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  js, err := json.Marshal(map[string]interface{}{
    "ref": ref_id,
  })
  HandleError(err)
  w.Write(js)
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
