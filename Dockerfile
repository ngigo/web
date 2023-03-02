FROM golang:1.16-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ngigo-web .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/ngigo-web .

CMD ["/app/ngigo-web"]