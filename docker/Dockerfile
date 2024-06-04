# build container
FROM golang:1.22.2-alpine AS builder
LABEL APP_VERSION="0.1.0"
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o chess-engine .

# app container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/chess-engine .
CMD ["./main"]
