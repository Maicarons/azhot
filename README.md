<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="Logo" width="128" height="128" />
</p>

<p align="center">
  <img src="banner.jpg" alt="Banner" style="max-width:100%;height:auto;" />
</p>

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![License](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

## 🌐 Traduções / Translations

- [简体中文](README.md)
- [English](README.en.md)
- [Français](README.fr.md)
- [한국어](README.ko.md)
- [Español](README.es.md)
- [Português](README.pt.md)

---





> 一个提供各大平台热搜API的聚合服务



## 📖 目录

- [项目简介](#项目简介)
- [功能特性](#功能特性)
- [支持平台](#支持平台)
- [快速开始](#快速开始)
- [API使用方法](#api使用方法)
- [MCP服务器](#mcp服务器)
- [开发贡献](#开发贡献)
- [许可证](#许可证)
- [问题反馈](#问题反馈)

## 项目简介

`azhot` 是一个聚合各大平台热搜数据的API服务，提供统一的接口访问各大平台的热搜内容。项目使用Go语言开发，基于Fiber框架构建，支持实时获取各大平台的热搜榜单数据。

## 功能特性

- 🚀 统一API接口，获取各大平台热搜数据
- ⚡ 高性能，使用`Go`+`Fiber v2`开发，带原生缓存机制 + 访问控制
- 🔄 定时更新热搜数据到数据库 【支持SQLite + MySQL + 可扩展其他DB】
- 📚 [Swagger API文档](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 RESTful API设计
- 📦 自带示例[前端](/frontend)
- 🔌 支持WebSocket实时数据推送
- 🤖 **新增** 支持AI Model Context Protocol (MCP) 服务器

## 项目结构
```
azhot/
├── all/                 # all功能代码
├── app/                 # 主程序代码
├── config/              # 读取配置文件
├── docs/                # swagger API文档
├── model/               # 数据库模型
├── mcp/                 # AI Model Context Protocol 服务器
├── router/              # 路由配置
├── service/             # 业务逻辑
├── websocket/           # WebSocket功能
├── frontend/            # 模板文件
├── .env                 # 环境变量
├── Dockerfile           # Docker构建文件
├── go.mod               # Go模块定义
├── main.go              # 主程序文件
└── README.md            # 项目说明文档
```

## 支持平台

| 名称 | 路由名 | 可用性 |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| 360搜索 | 360search | ✅ |
| AcFun | acfun | ✅ |
| 百度 | baidu | ✅ |
| 哔哩哔哩 | bilibili | ✅ |
| 央视网 | cctv | ✅ |
| CSDN | csdn | ✅ |
| 懂球帝 | dongqiudi | ✅ |
| 豆瓣 | douban | ✅ |
| 抖音 | douyin | ✅ |
| GitHub | github | ✅ |
| 国家地理 | guojiadili | ✅ |
| 历史上的今天 | historytoday | ✅ |
| 虎扑 | hupu | ✅ |
| IT之家 | ithome | ✅ |
| 梨视频 | lishipin | ✅ |
| 南方周末 | nanfang | ✅ |
| 澎湃新闻 | pengpai | ✅ |
| 腾讯新闻 | qqnews | ✅ |
| 夸克 | quark | ✅ |
| 人民网 | renmin | ✅ |
| 搜狗 | sougou | ✅ |
| 搜狐 | souhu | ✅ |
| 今日头条 | toutiao | ✅ |
| V2EX | v2ex | ✅ |
| 网易新闻 | wangyinews | ✅ |
| 微博 | weibo | ✅ |
| 新京报 | xinjingbao | ✅ |
| 知乎 | zhihu | ✅ |

## 快速开始

### 环境要求

- Go >= 1.18
- MySQL (可选，用于数据存储)

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置环境
```bash
# 复制配置文件
cp .env.example .env
# 编辑配置文件
vim .env
```

4. 生成API文档
```bash
swag init
```

5. 运行项目
```bash
# 开发模式运行
make dev

# 或者构建后运行
make run
```

### 使用Docker运行

```bash
# 构建镜像
docker build -t azhot .

# 运行容器
docker run -d -p 8080:8080 azhot
```

### 环境变量配置

项目使用 `.env` 文件进行配置，以下是可用的环境变量：

#### 服务器配置
- `SERVER_HOST`: 服务器主机地址，默认为 `localhost`
- `SERVER_PORT`: 服务器端口，默认为 `8080`
- `TLS_ENABLED`: 是否启用TLS/HTTPS，默认为 `false`
- `TLS_CERT_FILE`: TLS证书文件路径，当 `TLS_ENABLED` 为 `true` 时必须提供
- `TLS_KEY_FILE`: TLS私钥文件路径，当 `TLS_ENABLED` 为 `true` 时必须提供

#### 数据库配置
- `DB_TYPE`: 数据库类型，支持 `sqlite` 和 `mysql`，默认为 `sqlite`
- `MYSQL_DSN`: MySQL 数据库连接字符串，当 `DB_TYPE` 为 `mysql` 时生效

#### MCP 配置
- `MCP_STDIO_ENABLED`: 是否启用 STDIO MCP 服务器，默认为 `false`
- `MCP_HTTP_ENABLED`: 是否启用 HTTP MCP 服务器，默认为 `false`
- `MCP_PORT`: HTTP MCP 服务器端口，默认为 `8081`

#### 调试配置
- `DEBUG`: 是否启用调试模式，默认为 `false`

#### CORS 配置
- `CORS_ALLOW_ORIGINS`: 允许的跨域请求来源，多个来源用逗号分隔，默认为空表示允许所有来源（仅在生产环境中推荐设置具体来源）

## API使用方法

### HTTP API

#### 获取所有平台列表

```
GET /list
```

获取所有支持的平台信息。

#### 获取特定平台热搜

```
GET /{platform}
```

例如获取微博热搜：
```
GET /zhihu
```

### WebSocket API

项目支持WebSocket实时数据推送，提供与HTTP API相同的路由结构。

#### 通用WebSocket端点

```
ws://localhost:8080/ws
```

连接后可以发送消息来订阅或请求特定平台数据。

#### 特定平台WebSocket端点

```
ws://localhost:8080/ws/{platform}
```

例如连接百度热搜WebSocket：
```
ws://localhost:8080/ws/baidu
```

#### WebSocket消息格式

```json
{
  "type": "subscribe|request|ping",
  "source": "平台名称，如baidu、zhihu等",
  "data": {}
}
```

- `subscribe`: 订阅特定平台的实时数据
- `request`: 请求一次性数据
- `ping`: 心跳消息

#### WebSocket端点列表

- 通用端点: `ws://localhost:8080/ws`
- 百度: `ws://localhost:8080/ws/{platform}`
- 所有平台聚合: `ws://localhost:8080/ws/all`
- 平台列表: `ws://localhost:8080/ws/list`
- 历史查询API:
  - `ws://localhost:8080/ws/history/{source}` - 获取指定平台的所有历史数据
  - `ws://localhost:8080/ws/history/{source}/{date}` - 获取指定平台、日期的所有小时数据
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - 获取指定平台、日期和小时的历史数据
- 以及其他所有HTTP API对应的WebSocket端点

### API响应格式

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "2026新年贺词",
      "url": "https://www.zhihu.com/search?q=2026新年贺词"
    },
    // ...
    {
      "index": 12,
      "title": "东北网友发现「小鼻嘎」老鼠",
      "url": "https://www.zhihu.com/search?q=东北网友发现「小鼻嘎」老鼠"
    }
  ]
}
```

## MCP服务器

项目现在集成了AI Model Context Protocol (MCP) 服务器，允许AI模型和智能助手通过标准化的协议访问热搜数据。

### 功能特性

- **标准化工具接口**: 提供标准的MCP工具列表和执行接口
- **热搜数据访问**: 支持通过工具获取各平台热搜数据
- **历史数据查询**: 支持查询历史热搜数据
- **多种部署模式**: 支持HTTP和STDIO两种部署模式

### 启用MCP服务器

在 `.env` 文件中配置以下选项：

```env
MCP_STDIO_ENABLED=true      # 启用STDIO MCP服务器
MCP_HTTP_ENABLED=true       # 启用HTTP MCP服务器
MCP_PORT=8081               # HTTP MCP服务器端口
```

### MCP工具列表

- `get_hot_search`: 获取指定平台的热搜数据
- `get_all_hot_search`: 获取所有平台的热搜数据聚合
- `get_history_data`: 获取指定平台的历史热搜数据

### MCP端点

- `/mcp/tools` - 获取可用工具列表
- `/mcp/tool/execute` - 执行指定工具
- `/mcp/prompts` - 获取可用提示词列表
- `/mcp/ping` - 健康检查端点
- `/mcp/.well-known/mcp-info` - MCP服务器元数据

### 使用示例

通过HTTP调用MCP工具：
```bash
curl -X POST http://localhost:8080/mcp/tool/execute \
  -H "Content-Type: application/json" \
  -d '{
    "method": "tool/execute",
    "params": {
      "name": "get_hot_search",
      "arguments": {
        "platform": "zhihu"
      }
    },
    "id": "req-1",
    "jsonrpc": "2.0"
  }'
```

## 开发贡献

我们欢迎任何形式的贡献！如果您想为项目做出贡献，请按照以下步骤操作：

1. Fork 本项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 本地开发

```bash
# 运行测试
dev.sh # 使用Air作为热重启调试工具
```

## CMake构建系统

项目现在支持使用CMake进行构建，支持Windows和Linux平台。

### 构建命令

```bash
# 构建当前平台
mkdir build && cd build
cmake ..
cmake --build . --target build

# 运行
cmake --build . --target run

# 开发模式运行
cmake --build . --target dev

# 跨平台构建（预定义平台）
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# 跨平台构建（使用脚本）
# Linux/macOS:
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows:
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# 打包（为所有支持的平台创建zip包）
cmake --build . --target package

# 清理构建产物
cmake --build . --target azhot_clean

# 运行测试
cmake --build . --target test

# 运行所有测试
cmake --build . --target test-all

# 格式化代码
cmake --build . --target fmt

# 整理依赖
cmake --build . --target tidy

# 静态分析
cmake --build . --target staticcheck

# 构建CI版本（不生成swagger文档）
cmake --build . --target build-ci
```

## 许可证

本项目采用 AGPL-3.0 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 问题反馈

如果你在使用过程中遇到问题或有任何建议，欢迎提交 Issue 或 Pull Request。

- 🐛 [问题报告](https://github.com/maicarons/azhot/issues)
- ✨ [功能建议](https://github.com/maicarons/azhot/issues)

---

> 🌟 如果这个项目对你有帮助，请给我们一个 Star！这将是对我们最大的支持！