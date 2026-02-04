FROM golang:1.23-alpine3.20

RUN apk add --no-cache git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
