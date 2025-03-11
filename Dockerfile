FROM golang:1.23-alpine

RUN apk add build-base

RUN mkdir -p /go/src/github.com/dinizgab/golang-tests

WORKDIR /go/src/github.com/dinizgab/golang-tests
COPY . .

RUN go get -t -v -d ./...
