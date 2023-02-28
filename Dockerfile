FROM golang:1.16-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ngigo-web .

FROM alpine:latest

WORKDIR /app

COPY /app/ngigo-web .

CMD ['bash', '/app/ngigo-web']