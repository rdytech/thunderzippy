FROM golang:1.23-alpine3.20

RUN apk add --no-cache git ca-certificates

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o app .

CMD ["./app"]
