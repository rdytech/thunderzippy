# Thunderzippy

Secure streaming zipped download bundles.

A microservice that takes a list of remote files and provides a url that streams
a zipped bundle of all those files.

Concept based on an idea from
http://engineroom.teamwork.com/how-to-securely-provide-a-zip-download-of-a-s3-file-bundle/
then simplified to work as a simple RESTful API.

## Install and Run

```
docker-compose build
docker-compose up
```

### Environment Variables

`REDIS_ADDRESS` : Address with port of Redis host e.g. `127.0.0.1:6379`
`PORT` : Local port the HTTP server will bind to

## Use

### Create zip
```
$ curl -H "Content-Type: application/json" -X POST -d '[{"Filepath":"images/example.jpg","Url":"https://upload.wikimedia.org/wikipedia/mediawiki/a/a9/Example.jpg"}]' http://localhost:8080/zip/

{"ref":"z95h07"}
```

### Download zip
```
$ curl http://localhost:8080/zip/?ref=z95h07 -o "download.zip"
$ unzip -l download.zip
Archive:  download.zip
  Length     Date   Time    Name
 --------    ----   ----    ----
    61136  02-26-16 03:00   images/example.jpg
 --------                   -------
    61136                   1 file
```

### License

Thunderzippy is released under the [MIT License](http://www.opensource.org/licenses/MIT).

Patches, suggestions and comments are welcome.
