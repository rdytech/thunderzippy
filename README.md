# Zippy

https://golang.org/doc/install

## Install
cd $GOPATH
go install github.com/s01ipsist/zippy

## Run
$ $GOPATH/bin/zippy
2016/02/16 20:07:16 Get:  http://localhost:3000/images/CC-attribution.png
2016/02/16 20:07:16 GET /zip/ 11.712815ms

## Use

```
$ curl -v http://localhost:8080/zip/
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /zip/ HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Disposition: attachment; filename="test.zip"
< Content-Type: application/zip
< Date: Tue, 16 Feb 2016 08:24:16 GMT
< Content-Length: 668
<
CC-attribution.png���PNG

IHDRX���LTEo��*+)FHEikh������������L�1�tRNS@��fbKGD�H pHYs

                                                            ��mIDAT8ˍ�Ao�@`�>�JR�Wc-�Ҧ�Y�T�؄�e�Ҧ�~vҸ�.�\�2y�����u��U�ق�?�K:s�""o�""�90��9`��9�t�9�H <�;k9�%>L<FJrWm��X�.1"I~#V�,f��7���"_�mIr�8
\�
* Connection #0 to host localhost left intact
G�u�A��� ���Kp?����TC�q���Ip����h�Z�Ր���-���E@��}���g�_ׂCI��1nWsw?�b��]x̀��7�'����I�H��mE��Ty{���g�E�b,��o�_o^��pev  c���VZ㐨Q�esa��!�0<���)�������Tc�m���lt����r:����^��@*<�IEND�B`���P�䡂�䡂�CC-attribution.pngPK@F
```

http://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
Note that io.Copy reads 32kb (maximum) from input and writes them to output, then repeats. So don't worry about memory
