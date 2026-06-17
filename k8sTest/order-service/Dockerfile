FROM golang:1.24 AS builder
ENV CGO_ENABLED=0 GOOS=linux
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .
RUN go build -o order-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-service .
CMD ["./order-service"]