FROM golang:1.24 AS builder

ENV CGO_ENABLED=0 GOOS=linux
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build -o user-service ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
RUN chmod +x ./user-service

CMD ["./user-service"]