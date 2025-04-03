# Imagen base de Go
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./src

EXPOSE 3000
CMD ["./main"]
