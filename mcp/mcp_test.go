package mcp

import (
	"api/config"
	"api/service"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMCPHandler(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 验证处理器是否正确初始化
	assert.NotNil(t, handler)
	assert.Equal(t, service, handler.service)
	assert.Equal(t, config, handler.config)
	assert.NotNil(t, handler.tools)
	assert.NotEmpty(t, handler.tools)
}

func TestRegisterTools(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 验证工具是否已注册
	expectedTools := []string{"get_hot_search", "get_all_hot_search", "get_history_data"}
	for _, expectedTool := range expectedTools {
		_, exists := handler.tools[expectedTool]
		assert.True(t, exists, "Tool %s should be registered", expectedTool)
	}

	// 验证工具描述是否正确
	getHotSearchTool, exists := handler.tools["get_hot_search"]
	assert.True(t, exists)
	assert.Equal(t, "get_hot_search", getHotSearchTool.Name)
	assert.Contains(t, getHotSearchTool.Description, "获取各大平台的热搜数据")
}

func TestHandleListTools(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 调用handleListTools
	responseBytes, err := handler.handleListTools("test-id")
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "test-id", response.ID)
	assert.Equal(t, "2.0", response.Version)
	assert.NotNil(t, response.Result)

	// 验证结果包含工具列表
	resultMap, ok := response.Result.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, resultMap, "tools")

	tools, ok := resultMap["tools"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, tools)
}

func TestHandlePing(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 调用handlePing
	responseBytes, err := handler.handlePing("ping-id")
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "ping-id", response.ID)
	assert.Equal(t, "2.0", response.Version)
	assert.NotNil(t, response.Result)

	// 验证结果包含pong消息
	resultMap, ok := response.Result.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, resultMap, "message")
	assert.Equal(t, "pong", resultMap["message"])
}

func TestHandleListPrompts(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 调用handleListPrompts
	responseBytes, err := handler.handleListPrompts("prompts-id")
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "prompts-id", response.ID)
	assert.Equal(t, "2.0", response.Version)
	assert.NotNil(t, response.Result)

	// 验证结果包含提示列表
	resultMap, ok := response.Result.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, resultMap, "prompts")

	prompts, ok := resultMap["prompts"].([]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, prompts)
}

func TestGetAppFunctionByName(t *testing.T) {
	// 测试获取存在的函数
	fn := getAppFunctionByName("Baidu")
	assert.NotNil(t, fn, "Should be able to get Baidu function")

	// 测试获取不存在的函数
	nilFn := getAppFunctionByName("NonExistentFunction")
	assert.Nil(t, nilFn, "Should return nil for non-existent function")
}

func TestCreateErrorResponse(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 调用createErrorResponse
	responseBytes, err := handler.createErrorResponse("error-id", 500, "Test error message")
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "error-id", response.ID)
	assert.Equal(t, "2.0", response.Version)
	assert.NotNil(t, response.Error)
	assert.Equal(t, 500, response.Error.Code)
	assert.Equal(t, "Test error message", response.Error.Message)
}

func TestHandleRequestWithInvalidJSON(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 尝试处理无效的JSON
	invalidJSON := []byte("{ invalid json }")
	responseBytes, err := handler.HandleRequest(invalidJSON)
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "", response.ID)
	assert.NotNil(t, response.Error)
	assert.Equal(t, -32700, response.Error.Code)
	assert.Contains(t, response.Error.Message, "Parse error")
}

func TestHandleRequestWithUnknownMethod(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 创建未知方法的请求
	request := Request{
		Method:  "unknown/method",
		ID:      "unknown-id",
		Version: "2.0",
	}
	requestBytes, _ := json.Marshal(request)

	responseBytes, err := handler.HandleRequest(requestBytes)
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "unknown-id", response.ID)
	assert.NotNil(t, response.Error)
	assert.Equal(t, -32601, response.Error.Code)
	assert.Contains(t, response.Error.Message, "Method not found")
}

func TestExecuteGetAllHotSearch(t *testing.T) {
	// 创建服务和配置
	service := &service.HotSearchService{}
	config := &config.Config{}

	// 创建MCP处理器
	handler := NewMCPHandler(service, config)

	// 调用executeGetAllHotSearch
	responseBytes, err := handler.executeGetAllHotSearch("all-id")
	assert.NoError(t, err)

	// 解析响应
	var response Response
	err = json.Unmarshal(responseBytes, &response)
	assert.NoError(t, err)

	assert.Equal(t, "all-id", response.ID)
	assert.Equal(t, "2.0", response.Version)
	assert.NotNil(t, response.Result)
}
