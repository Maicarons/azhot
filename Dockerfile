# 使用多阶段构建以减小最终镜像大小
# 第一阶段：构建应用
FROM golang:1.24-alpine AS builder

# 安装构建时依赖（如SQLite支持所需的包）
RUN apk add --no-cache git gcc musl-dev

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 生成API文档
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

# 构建应用
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# 第二阶段：创建运行时镜像
FROM alpine:latest

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN addgroup -g 65532 nonroot &&\
    adduser -u 65532 -G nonroot -s /bin/sh -D nonroot

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .

# 确保二进制文件具有可执行权限
RUN chmod +x ./main

# 创建数据库目录（如果使用SQLite）
RUN mkdir -p /data && chown -R nonroot:nonroot /data

# 暴露端口（默认8080，可通过环境变量配置）
EXPOSE 8080

# 使用非root用户运行
USER nonroot

# 启动应用
CMD ["./main"]