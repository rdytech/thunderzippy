# Zippy

https://golang.org/doc/install

## Install
cd $GOPATH
go install github.com/s01ipsist/zippy

## Run
$ $GOPATH/bin/zippy
2016/02/17 10:55:20 Thunderzippy is go
2016/02/17 10:55:23 GET     /zip/
2016/02/17 10:55:23 adding: 404 http://localhost:3000/images/CC-attribution-not-found.png
2016/02/17 10:55:23 adding: 200 http://localhost:3000/images/facebook-small.png
2016/02/17 10:55:23 Tzipped:  (68.016689ms)

## Use

```
$ curl http://localhost:8080/zip/ -o "test.zip"
$ unzip -l test.zip
Archive:  test.zip
  Length     Date   Time    Name
 --------    ----   ----    ----
     1042  02-16-16 21:55   images/facebook-small.png
 --------                   -------
     1042                   1 file
```

http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
Note that io.Copy reads 32kb (maximum) from input and writes them to output, then repeats. So don't worry about memory
