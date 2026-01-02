<div style="text-align: center;">

# azhot

<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/3199/3199306.png" alt="Logo" width="128" height="128" />
</p>

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![License](https://img.shields.io/github/license/maicarons/azhot)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/maicarons/azhot)](https://goreportcard.com/report/github.com/maicarons/azhot)

</div>

> A聚合service that provides hot search APIs for major platforms

## 📖 Table of Contents

- [Project Introduction](#project-introduction)
- [Features](#features)
- [Supported Platforms](#supported-platforms)
- [Quick Start](#quick-start)
- [API Usage](#api-usage)
- [MCP Server](#mcp-server)
- [Development and Contribution](#development-and-contribution)
- [License](#license)
- [Issue Feedback](#issue-feedback)

## Project Introduction

`azhot` is an API service that aggregates hot search data from major platforms, providing a unified interface to access hot search content from various platforms. The project is developed in Go language and built based on the Fiber framework, supporting real-time retrieval of hot search ranking data from major platforms.

## Features

- 🚀 Unified API interface to retrieve hot search data from major platforms
- ⚡ High performance, developed with `Go`+`Fiber v2`, with native caching mechanism + access control
- 🔄 Scheduled updates of hot search data to database [Supports SQLite + MySQL + Extensible other DBs]
- 📚 [Swagger API Documentation](https://github.com/maicarons/azhot/blob/main/docs/swagger.yaml)
- 🌐 RESTful API design
- 📦 Includes example [Frontend](/frontend)
- 🔌 Supports WebSocket real-time data push
- 🤖 **New** Supports AI Model Context Protocol (MCP) server

## Project Structure
```
azhot/
├── all/                 # All functionality code
├── app/                 # Main program code
├── config/              # Configuration file reading
├── docs/                # Swagger API documentation
├── model/               # Database models
├── mcp/                 # AI Model Context Protocol server
├── router/              # Routing configuration
├── service/             # Business logic
├── websocket/           # WebSocket functionality
├── frontend/            # Template files
├── .env                 # Environment variables
├── Dockerfile           # Docker build file
├── go.mod               # Go module definition
├── main.go              # Main program file
└── README.md            # Project documentation
```

## Supported Platforms

| Name | Route Name | Availability |
|:----:|:------:|:------:|
| 360doc | 360doc | ✅ |
| 360 Search | 360search | ✅ |
| AcFun | acfun | ✅ |
| Baidu | baidu | ✅ |
| Bilibili | bilibili | ✅ |
| CCTV | cctv | ✅ |
| CSDN | csdn | ✅ |
| Dongqiudi | dongqiudi | ✅ |
| Douban | douban | ✅ |
| Douyin | douyin | ✅ |
| GitHub | github | ✅ |
| National Geographic | guojiadili | ✅ |
| History Today | historytoday | ✅ |
| Hupu | hupu | ✅ |
| IT Home | ithome | ✅ |
| Pear Video | lishipin | ✅ |
| Southern Weekly | nanfang | ✅ |
| Pengpai News | pengpai | ✅ |
| Tencent News | qqnews | ✅ |
| Quark | quark | ✅ |
| People's Daily Online | renmin | ✅ |
| Sogou | sougou | ✅ |
| Sohu | souhu | ✅ |
| Toutiao | toutiao | ✅ |
| V2EX | v2ex | ✅ |
| NetEase News | wangyinews | ✅ |
| Weibo | weibo | ✅ |
| Xinjing Daily | xinjingbao | ✅ |
| Zhihu | zhihu | ✅ |

## Quick Start

### Environment Requirements

- Go >= 1.18
- MySQL (Optional, for data storage)

### Installation Steps

1. Clone the project
```bash
git clone https://github.com/maicarons/azhot.git
cd azhot
```

2. Install dependencies
```bash
go mod tidy
```

3. Configure environment
```bash
# Copy configuration file
cp .env.example .env
# Edit configuration file
vim .env
```

4. Generate API documentation
```bash
swag init
```

5. Run the project
```bash
# Run in development mode
make dev

# Or build and run
make run
```

### Running with Docker

```bash
# Build image
docker build -t azhot .

# Run container
docker run -d -p 8080:8080 azhot
```

## API Usage

### HTTP API

#### Get All Platform List

```
GET /list
```

Retrieve information for all supported platforms.

#### Get Hot Search for Specific Platform

```
GET /{platform}
```

For example, to get Zhihu hot search:
```
GET /zhihu
```

### WebSocket API

The project supports WebSocket real-time data push, providing the same routing structure as the HTTP API.

#### General WebSocket Endpoint

```
ws://localhost:8080/ws
```

After connecting, you can send messages to subscribe to or request specific platform data.

#### Specific Platform WebSocket Endpoint

```
ws://localhost:8080/ws/{platform}
```

For example, connecting to Baidu hot search WebSocket:
```
ws://localhost:8080/ws/baidu
```

#### WebSocket Message Format

```json
{
  "type": "subscribe|request|ping",
  "source": "Platform name, such as baidu, zhihu, etc.",
  "data": {}
}
```

- `subscribe`: Subscribe to real-time data for a specific platform
- `request`: Request one-time data
- `ping`: Heartbeat message

#### WebSocket Endpoints List

- General endpoint: `ws://localhost:8080/ws`
- Baidu: `ws://localhost:8080/ws/{platform}`
- All platforms aggregation: `ws://localhost:8080/ws/all`
- Platform list: `ws://localhost:8080/ws/list`
- Historical query API:
  - `ws://localhost:8080/ws/history/{source}` - Get all historical data for a specified platform
  - `ws://localhost:8080/ws/history/{source}/{date}` - Get all hourly data for a specified platform and date
  - `ws://localhost:8080/ws/history/{source}/{date}/{hour}` - Get historical data for a specified platform, date, and hour
- And all other WebSocket endpoints corresponding to HTTP APIs

### API Response Format

```json
{
  "code": 200,
  "icon": "https://static.zhihu.com/static/favicon.ico",
  "message": "zhihu",
  "obj": [
    {
      "index": 1,
      "title": "2026 New Year Greetings",
      "url": "https://www.zhihu.com/search?q=2026 New Year Greetings"
    },
    // ...
    {
      "index": 12,
      "title": "Northeast netizens discover 'Xiao Biga' mouse",
      "url": "https://www.zhihu.com/search?q=Northeast netizens discover 'Xiao Biga' mouse"
    }
  ]
}
```

## MCP Server

The project now integrates an AI Model Context Protocol (MCP) server, allowing AI models and intelligent assistants to access hot search data through a standardized protocol.

### Features

- **Standardized Tool Interface**: Provides standard MCP tool list and execution interface
- **Hot Search Data Access**: Supports retrieving hot search data for each platform through tools
- **Historical Data Query**: Supports querying historical hot search data
- **Multiple Deployment Modes**: Supports both HTTP and STDIO deployment modes

### Enabling MCP Server

Configure the following options in the `.env` file:

```env
MCP_STDIO_ENABLED=true      # Enable STDIO MCP server
MCP_HTTP_ENABLED=true       # Enable HTTP MCP server
MCP_PORT=8081               # HTTP MCP server port
```

### MCP Tools List

- `get_hot_search`: Get hot search data for a specified platform
- `get_all_hot_search`: Get aggregated hot search data for all platforms
- `get_history_data`: Get historical hot search data for a specified platform

### MCP Endpoints

- `/mcp/tools` - Get list of available tools
- `/mcp/tool/execute` - Execute specified tool
- `/mcp/prompts` - Get list of available prompts
- `/mcp/ping` - Health check endpoint
- `/mcp/.well-known/mcp-info` - MCP server metadata

### Usage Example

Calling MCP tool via HTTP:
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

For more details, please refer to the [MCP Server Documentation](mcp/README.md).

## Development and Contribution

We welcome any form of contribution! If you'd like to contribute to the project, please follow these steps:

1. Fork this project
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Create a Pull Request

### Local Development

```bash
# Run tests
dev.sh # Using Air as hot-reload debugging tool
```

## CMake Build System

The project now supports building with CMake, supporting both Windows and Linux platforms.

### Build Commands

```bash
# Build for current platform
mkdir build && cd build
cmake ..
cmake --build . --target build

# Run
cmake --build . --target run

# Run in development mode
cmake --build . --target dev

# Cross-platform build (predefined platforms)
cmake --build . --target build-platform-linux
cmake --build . --target build-platform-windows
cmake --build . --target build-platform-darwin
cmake --build . --target build-platform-linux-arm64
cmake --build . --target build-platform-windows-arm64

# Cross-platform build (using script)
# Linux/macOS:
./build_platform.sh linux
./build_platform.sh windows
./build_platform.sh darwin

# Windows:
build_platform.bat linux
build_platform.bat windows
build_platform.bat darwin

# Package (create zip packages for all supported platforms)
cmake --build . --target package

# Clean build artifacts
cmake --build . --target azhot_clean

# Run tests
cmake --build . --target test

# Run all tests
cmake --build . --target test-all

# Format code
cmake --build . --target fmt

# Tidy dependencies
cmake --build . --target tidy

# Static analysis
cmake --build . --target staticcheck

# Build CI version (without generating swagger docs)
cmake --build . --target build-ci
```

## License

This project is licensed under the AGPL-3.0 License - see the [LICENSE](LICENSE) file for details.

## Issue Feedback

If you encounter any problems or have suggestions while using the project, feel free to submit an Issue or Pull Request.

- 🐛 [Issue Report](https://github.com/maicarons/azhot/issues)
- ✨ [Feature Request](https://github.com/maicarons/azhot/issues)

---

> 🌟 If this project is helpful to you, please give us a Star! This would be the greatest support for us!