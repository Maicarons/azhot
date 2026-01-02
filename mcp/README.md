# AI Model Context Protocol (MCP) 服务器集成

## 概述

本项目现在集成了AI Model Context Protocol (MCP) 服务器，允许AI模型和智能助手通过标准化的协议访问热搜数据。MCP是一种开放协议，为AI模型提供了一种标准化的方式与外部工具、API和数据源进行交互。

## 功能特性

### 1. 标准化工具接口
- 提供标准化的工具列表 (`tools/list`)
- 支持工具执行 (`tool/execute`)
- 提供提示词管理 (`prompts/list`)

### 2. 热搜数据访问工具
MCP服务器提供了以下工具来访问热搜数据：

#### `get_hot_search`
- **描述**: 获取指定平台的热搜数据
- **参数**:
  - `platform` (string, required): 平台名称 (如: baidu, bilibili, zhihu, weibo, 等)
- **示例**:
  ```json
  {
    "method": "tool/execute",
    "params": {
      "name": "get_hot_search",
      "arguments": {
        "platform": "baidu"
      }
    },
    "id": "req-1",
    "jsonrpc": "2.0"
  }
  ```

#### `get_all_hot_search`
- **描述**: 获取所有平台的热搜数据聚合
- **参数**: 无
- **示例**:
  ```json
  {
    "method": "tool/execute",
    "params": {
      "name": "get_all_hot_search"
    },
    "id": "req-2",
    "jsonrpc": "2.0"
  }
  ```

#### `get_history_data`
- **描述**: 获取指定平台的历史热搜数据
- **参数**:
  - `platform` (string, required): 平台名称
  - `date` (string, required): 日期 (格式: YYYY-MM-DD)
  - `hour` (string, optional): 小时 (格式: HH)
- **示例**:
  ```json
  {
    "method": "tool/execute",
    "params": {
      "name": "get_history_data",
      "arguments": {
        "platform": "baidu",
        "date": "2025-01-01",
        "hour": "12"
      }
    },
    "id": "req-3",
    "jsonrpc": "2.0"
  }
  ```

### 3. 提示词功能
- `analyze_hot_search_trends`: 分析当前热搜趋势，识别热门话题和用户兴趣
- `compare_platform_topics`: 比较不同平台的热门话题，分析差异和共同点

## 配置选项

MCP服务器支持以下配置选项，可以通过环境变量进行配置：

- `MCP_STDIO_ENABLED`: 启用/禁用STDIO MCP服务器 (true/false, 默认: false)
- `MCP_HTTP_ENABLED`: 启用/禁用HTTP MCP服务器 (true/false, 默认: false)
- `MCP_PORT`: HTTP MCP服务器端口 (默认: 8081)

### 示例配置 (.env)
```env
MCP_STDIO_ENABLED=true
MCP_HTTP_ENABLED=true
MCP_PORT=8081
```

## 部署方式

### 1. HTTP服务器模式
当 `MCP_HTTP_ENABLED=true` 时，MCP服务器将在指定端口上启动HTTP服务。

- 工具列表端点: `GET /mcp/tools`
- 工具执行端点: `POST /mcp/tool/execute`
- 提示列表端点: `GET /mcp/prompts`
- Ping端点: `GET /mcp/ping`
- 发现端点: `GET /mcp/.well-known/mcp-info`

### 2. STDIO模式
当 `MCP_STDIO_ENABLED=true` 时，MCP服务器将通过标准输入输出与AI助手通信，遵循JSON-RPC 2.0协议。

## API端点

除了MCP协议端点外，项目还提供了以下MCP相关端点：

- `/mcp/tools` - 获取可用工具列表
- `/mcp/tool/execute` - 执行指定工具
- `/mcp/prompts` - 获取可用提示词列表
- `/mcp/ping` - 健康检查端点
- `/mcp/.well-known/mcp-info` - MCP服务器元数据

## 使用示例

### 通过HTTP调用工具
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

### 通过STDIO使用
当STDIO模式启用时，MCP服务器将读取标准输入中的JSON-RPC请求并输出响应。

## 支持的平台

MCP服务器支持以下平台的热搜数据获取：
- Baidu (百度)
- Zhihu (知乎)
- Weibo (微博)
- Bilibili (B站)
- Douyin (抖音)
- Toutiao (今日头条)
- V2EX
- GitHub
- 以及更多平台...

## 故障排除

1. **MCP服务器未启动**: 确认环境变量已正确设置
2. **工具执行失败**: 检查参数格式是否正确
3. **连接问题**: 确认端口和网络配置

## 安全考虑

- MCP服务器遵循标准协议，确保通信安全
- 所有API调用受原有系统安全机制保护
- 建议在生产环境中使用适当的访问控制

## 扩展性

MCP服务器设计为可扩展，可以轻松添加新工具和功能，以支持更多的数据源和操作。