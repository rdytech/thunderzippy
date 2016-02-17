# Zippy

https://golang.org/doc/install

## Install
cd $GOPATH
go install github.com/s01ipsist/zippy

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
$ $GOPATH/bin/zippy

2016/02/17 16:32:27 Thunderzippy is go
2016/02/17 16:32:35 GET   /zip/?ref=1
2016/02/17 16:32:35 adding: 404 http://localhost:3000/images/CC-attribution-not-found.png
2016/02/17 16:32:35 adding: 200 http://localhost:3000/images/facebook-small.png
2016/02/17 16:32:35 Thunderzipped:  2 files (43.065781ms)

## Use

```
$ curl http://localhost:8080/zip/?ref=1 -o "test.zip"
$ unzip -l test.zip
Archive:  test.zip
  Length     Date   Time    Name
 --------    ----   ----    ----
     1042  02-17-16 03:32   images/facebook-small.png
 --------                   -------
     1042                   1 file
```

http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
Note that io.Copy reads 32kb (maximum) from input and writes them to output, then repeats. So don't worry about memory
