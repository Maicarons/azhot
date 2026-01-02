package websocket

import (
	"api/config"
	"api/service"
	"encoding/json"
	"testing"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestWebSocketManager 测试WebSocket管理器
func TestWebSocketManager(t *testing.T) {
	// 创建配置
	cfg := &config.Config{
		Debug: true,
	}

	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建Fiber应用
	app := fiber.New()

	// 设置WebSocket路由
	SetupWebSocketRoutes(app, hotSearchService, cfg)

	// 这里我们不实际启动服务器，而是验证代码结构
	assert.NotNil(t, app)
}

// TestMessageStructure 测试消息结构
func TestMessageStructure(t *testing.T) {
	msg := Message{
		Type:   "response",
		Source: "baidu",
		Data:   "test data",
	}

	jsonData, err := json.Marshal(msg)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "response")
	assert.Contains(t, string(jsonData), "baidu")
}

// TestClientStructure 测试客户端结构
func TestClientStructure(t *testing.T) {
	// 创建一个模拟的WebSocket连接（仅用于测试结构）
	var conn *websocket.Conn
	client := &Client{
		Conn: conn,
		Type: "baidu",
	}

	assert.Equal(t, "baidu", client.Type)
}
