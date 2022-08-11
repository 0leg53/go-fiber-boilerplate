FROM golang:1.18.2-alpine3.15

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY ./.env /app
