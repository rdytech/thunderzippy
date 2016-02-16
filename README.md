# Zippy

https://golang.org/doc/install

## Install
cd $GOPATH
go install github.com/s01ipsist/zippy

## Run
$GOPATH/bin/zippy

## Use

```
$ curl -v http://localhost:8080/
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Tue, 16 Feb 2016 06:07:46 GMT
< Content-Length: 19
<
404 page not found
```

```
$ curl -v http://localhost:8080/hi/you
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hi/you HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 16 Feb 2016 06:20:47 GMT
< Content-Length: 21
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hi there, I love you!
```
