FROM golang:1.20.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy