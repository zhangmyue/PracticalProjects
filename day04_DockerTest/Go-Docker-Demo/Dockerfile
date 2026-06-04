# 1. 使用官方 Go 镜像作为构建环境
FROM golang:1.22-alpine AS builder

# 2. 设置工作目录（镜像内部）
WORKDIR /app

# 3. 复制 go.mod 并下载依赖
COPY go.mod .
RUN go mod download

# 4. 复制源代码
COPY . .

# 5. 编译 Go 程序（生成可执行文件 app）
RUN go build -o app .

# -------------------- 第二阶段：超小体积运行镜像 --------------------
FROM alpine:latest
WORKDIR /app

# 复制编译好的可执行文件
COPY --from=builder /app/app .

# 声明容器内部端口
EXPOSE 8080

# 启动程序
CMD ["./app"]