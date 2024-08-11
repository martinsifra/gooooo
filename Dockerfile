# Based on https://docs.docker.com/language/golang/build-images/

FROM golang:1.22.6-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /gooooo

CMD ["/gooooo"]