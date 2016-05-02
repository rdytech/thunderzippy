# Thunderzippy

Secure streaming zipped download bundles.

A microservice that takes a list of remote files and provides a url that streams
a zipped bundle of all those files.

Concept based on an idea from
http://engineroom.teamwork.com/how-to-securely-provide-a-zip-download-of-a-s3-file-bundle/

## Install
https://golang.org/doc/install
```
cd $GOPATH
# this is a workaround for private git repos, standard would be
# go get github.com/jobready/thunderzippy
mkdir -p src/github.com/jobready
cd src/github.com/jobready
git clone git@github.com:jobready/thunderzippy.git

go install github.com/jobready/thunderzippy
```

## Setup download data on Redis

example using Ruby redis client
```ruby
require 'json'
require 'redis'
r = Redis.new
files = [
  {
    url: "http://localhost:3000/images/CC-attribution-not-found.png",
    filepath: "images/CC-attribution.png"
  },{
    url: "http://localhost:3000/images/facebook-small.png",
    filepath: 'images/facebook-small.png'
  }
].to_json
r.set('zip:1', files)
```
## Configure

```
cp sample_conf.json $GOPATH/thunderzippy_conf.json
```

## Run
```
$ $GOPATH/bin/thunderzippy

2016/02/17 16:32:27 Thunderzippy is go
2016/02/17 16:32:35 GET   /zip/?ref=1
2016/02/17 16:32:35 adding: 404 http://localhost:3000/images/CC-attribution-not-found.png
2016/02/17 16:32:35 adding: 200 http://localhost:3000/images/facebook-small.png
2016/02/17 16:32:35 Thunderzipped:  2 files (43.065781ms)
```

## Use

### Create zip reference
```
$ curl -H "Content-Type: application/json" -X POST -d '[{"Filepath":"images/example.jpg","Url":"https://upload.wikimedia.org/wikipedia/mediawiki/a/a9/Example.jpg"}]' http://localhost:8080/zip/

{"ref":"z95h07"}
```

### Download zip by ID
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

http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
Note that io.Copy reads 32kb (maximum) from input and writes them to output, then repeats. So don't worry about memory


### Server installation
Thunderzippy can be run directly on port 80 or as a proxied service behind nginx using the sample_nginx.conf config.

sample_upstart.conf provides a sample Upstart script.
