# Stage 1: Build the Go binary
FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o inote-app ./cmd/

# Stage 2: Create minimal final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/inote-app .

EXPOSE 8080

CMD ["./inote-app"]
