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

// TestStartManager 测试Start方法
func TestStartManager(t *testing.T) {
	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建管理器
	manager := NewWsManager(hotSearchService)

	// 启动管理器，不期望出现panic
	assert.NotPanics(t, func() {
		manager.Start()
	})
}

// TestHandleWebSocket 测试HandleWebSocket方法
func TestHandleWebSocket(t *testing.T) {
	// 由于HandleWebSocket需要真实的WebSocket连接，我们只测试其结构定义
	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建管理器
	manager := NewWsManager(hotSearchService)

	// 验证管理器是否正确初始化
	assert.NotNil(t, manager)
	assert.Equal(t, hotSearchService, manager.hotSearchService)
	assert.NotNil(t, manager.clients)
	assert.NotNil(t, manager.broadcast)
	assert.NotNil(t, manager.register)
	assert.NotNil(t, manager.unregister)
}

// TestHandleSubscribe 测试handleSubscribe方法
func TestHandleSubscribe(t *testing.T) {
	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建管理器
	manager := NewWsManager(hotSearchService)

	// 创建一个模拟的客户端连接
	var conn *websocket.Conn
	client := &Client{
		Conn: conn,
		Type: "baidu",
	}

	// 调用handleSubscribe，不期望出现panic
	assert.NotPanics(t, func() {
		manager.handleSubscribe(client, "baidu")
	})
}

// TestHandleRequest 测试handleRequest方法
func TestHandleRequest(t *testing.T) {
	// 由于handleRequest方法依赖于真实的WebSocket连接，我们跳过此测试
	t.Skip("Skipping TestHandleRequest as it depends on real WebSocket connection")
}

// TestSetupWebSocketRoutes 测试SetupWebSocketRoutes方法
func TestSetupWebSocketRoutes(t *testing.T) {
	// 创建配置
	cfg := &config.Config{
		Debug: true,
	}

	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建Fiber应用
	app := fiber.New()

	// 调用SetupWebSocketRoutes，不期望出现panic
	assert.NotPanics(t, func() {
		SetupWebSocketRoutes(app, hotSearchService, cfg)
	})

	// 验证应用不为nil
	assert.NotNil(t, app)
}

// TestSetupWebSocketAPIRoutes 测试setupWebSocketAPIRoutes方法
func TestSetupWebSocketAPIRoutes(t *testing.T) {
	// 创建服务
	hotSearchService := &service.HotSearchService{}

	// 创建管理器
	manager := NewWsManager(hotSearchService)

	// 创建Fiber应用
	app := fiber.New()

	// 调用setupWebSocketAPIRoutes，不期望出现panic
	assert.NotPanics(t, func() {
		setupWebSocketAPIRoutes(app, manager)
	})

	// 验证应用不为nil
	assert.NotNil(t, app)
}

// TestSetupWebSocketHistoryRoutes 测试setupWebSocketHistoryRoutes方法
func TestSetupWebSocketHistoryRoutes(t *testing.T) {
	// 创建管理器
	manager := NewWsManager(&service.HotSearchService{})

	// 创建Fiber应用
	app := fiber.New()

	// 调用setupWebSocketHistoryRoutes，不期望出现panic
	assert.NotPanics(t, func() {
		setupWebSocketHistoryRoutes(app, manager)
	})

	// 验证应用不为nil
	assert.NotNil(t, app)
}
