# Thunderzippy

Secure streaming zipped download bundles.

A microservice that takes a list of remote files and provides a url that streams
a zipped bundle of all those files.

Concept based on an idea from
http://engineroom.teamwork.com/how-to-securely-provide-a-zip-download-of-a-s3-file-bundle/
then simplified to work as a simple RESTful API.

## Install
https://golang.org/doc/install
```
cd $GOPATH
go get github.com/jobready/thunderzippy
```

## Configure
```
cp sample_conf.json $GOPATH/thunderzippy_conf.json
```

This contains two options:

`RedisServerAndPort` : IP and port of the attached Redis resource

`Port` : the port that the HTTP server will bind to

## Run
```
$ $GOPATH/bin/thunderzippy
```

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

### Server installation
Thunderzippy can be run directly on port 80 or as a proxied service behind nginx using the sample_nginx.conf config.

sample_upstart.conf provides a sample Upstart script.

### Docker

```
cp docker_conf.json thunderzippy_conf.json
docker-compose build
docker-compose up
```

### License

Thunderzippy is released under the [MIT License](http://www.opensource.org/licenses/MIT).

Patches, suggestions and comments are welcome.
