# Docker 部署说明

<p align="center">
  <img src="banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

## 概述

本项目支持使用 Docker 容器化部署，提供了灵活的配置选项以适应不同的部署需求。

## 快速开始

### 使用 SQLite（默认）

```bash
# 构建并启动服务
docker-compose up -d

# 访问应用
# API 文档: http://localhost:8080/swagger
# 应用主页: http://localhost:8080
```

### 使用 MySQL

```bash
# 使用 MySQL 配置启动服务
docker-compose -f docker-compose.mysql.yml up -d

# 访问应用
# API 文档: http://localhost:8080/swagger
# 应用主页: http://localhost:8080
```

## Docker 镜像构建

### 构建镜像

```bash
# 使用默认 Dockerfile 构建
docker build -t azhot .

# 或者指定标签
docker build -t azhot:latest .
```

### 运行容器

```bash
# 运行容器（使用 SQLite）
docker run -p 8080:8080 -e DB_TYPE=sqlite -v azhot_data:/data azhot

# 运行容器（使用 MySQL）
docker run -p 8080:8080 \
  -e DB_TYPE=mysql \
  -e MYSQL_DSN=user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local \
  azhot
```

## 环境变量配置

| 环境变量 | 默认值 | 说明 |
|---------|--------|------|
| SERVER_HOST | localhost | 服务器主机地址 |
| SERVER_PORT | 8080 | 服务器端口 |
| DB_TYPE | sqlite | 数据库类型 (sqlite 或 mysql) |
| MYSQL_DSN | - | MySQL 数据库连接字符串 (当 DB_TYPE=mysql 时使用) |
| DEBUG | false | 是否启用调试模式 |

## 持久化数据

- **SQLite**: 数据库文件存储在 `/data/hot_search.db`，建议挂载卷以持久化数据
- **MySQL**: 使用命名卷 `mysql_data` 来持久化数据

## 部署最佳实践

### 生产环境部署

```yaml
# docker-compose.prod.yml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "80:8080"  # 映射到标准HTTP端口
    environment:
      - SERVER_HOST=0.0.0.0
      - SERVER_PORT=8080
      - DB_TYPE=sqlite
      - DEBUG=false
    volumes:
      - ./data:/data
      - ./logs:/app/logs  # 挂载日志目录
    restart: unless-stopped
    deploy:
      resources:
        limits:
          memory: 512M
          cpus: '0.5'
```

### 使用反向代理

```yaml
# docker-compose.proxy.yml
version: '3.8'

services:
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
    depends_on:
      - app

  app:
    build: .
    environment:
      - SERVER_HOST=0.0.0.0
      - SERVER_PORT=8080
      - DB_TYPE=sqlite
    volumes:
      - ./data:/data
    restart: unless-stopped
```

## 常见问题

### 1. 数据库权限问题

确保挂载的数据卷具有正确的权限：

```bash
# 创建数据目录并设置权限
mkdir -p data
chown -R 65532:65532 data  # nonroot 用户ID
```

### 2. 构建时依赖问题

如果构建失败，请确保：

- Docker 版本 >= 20.10
- 网络连接正常
- 有足够的磁盘空间

### 3. 运行时错误

检查容器日志：

```bash
# 查看应用日志
docker-compose logs app

# 查看 MySQL 日志
docker-compose logs db
```

## 安全考虑

- 避免在生产环境中使用默认的数据库密码
- 使用非 root 用户运行容器
- 定期更新基础镜像
- 限制容器资源使用

## 监控和日志

应用运行时会输出日志到标准输出，可通过以下方式查看：

```bash
# 实时查看日志
docker-compose logs -f app

# 查看特定时间的日志
docker-compose logs --since "1h" app
```

## 更新应用

```bash
# 拉取最新代码
git pull

# 重新构建镜像
docker-compose build

# 重启服务
docker-compose up -d
```