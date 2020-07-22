ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /go/src/ingrid-app

COPY . /go/src/ingrid-app

RUN go mod download
RUN go build -o ingrid-app

EXPOSE 8080

ENTRYPOINT ["./ingrid-app"]