## Builder
FROM golang:latest

RUN mkdir -p /app

WORKDIR /app
COPY    . .
RUN     go build src/main.go

EXPOSE 3000