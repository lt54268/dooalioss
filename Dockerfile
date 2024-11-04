# 第一阶段：构建阶段
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装依赖工具
RUN apk add --no-cache git

# 将 go.mod 和 go.sum 文件复制到工作目录
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 编译项目
RUN go build -o dooalioss-app .

# 第二阶段：运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 复制可执行文件和其他需要的文件
COPY --from=builder /app/dooalioss-app .

# 如果你的项目使用.env文件
COPY --from=builder /app/.env . 

# 暴露应用端口
EXPOSE 3030

# 启动应用程序
CMD ["./dooalioss-app"]
