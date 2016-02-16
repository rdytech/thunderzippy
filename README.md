# Zippy

https://golang.org/doc/install

## Install
cd $GOPATH
go install github.com/s01ipsist/zippy

## Run
$ $GOPATH/bin/zippy
2016/02/17 10:29:33 Get:  http://localhost:3000/images/CC-attribution.png
2016/02/17 10:29:33 Get:  http://localhost:3000/images/facebook-small.png
2016/02/17 10:29:33 GET /zip/ 61.397018ms

## Use

```
$ curl http://localhost:8080/zip/ -o "test.zip"
$ unzip -l test.zip
Archive:  test.zip
  Length     Date   Time    Name
 --------    ----   ----    ----
      647  02-16-16 21:27   images/CC-attribution.png
     1042  02-16-16 21:27   images/facebook-small.png
 --------                   -------
     1689                   2 files
```

http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
Note that io.Copy reads 32kb (maximum) from input and writes them to output, then repeats. So don't worry about memory
