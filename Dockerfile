FROM golang:1.22-alpine as base

WORKDIR /app

RUN apk update && \
  apk upgrade && \
  apk add --no-cache \
  libc-dev \
  make

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.22.0

COPY go.mod go.sum ./
RUN go mod download 

COPY .air.toml .
