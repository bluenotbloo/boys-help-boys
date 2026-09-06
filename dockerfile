FROM golang:1.20-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod tidy
RUN go build -o boys-help-boys

