# Builds stage
FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/bannerlord/main.go

EXPOSE 8080
CMD ["/app/main"]