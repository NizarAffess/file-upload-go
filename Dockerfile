FROM golang:1.16-alpine AS build

WORKDIR /go/src/go-file-upload

COPY  go.mod ./
COPY  go.sum ./

RUN go mod download

COPY . .

RUN go get  github.com/pilu/fresh

ENTRYPOINT fresh