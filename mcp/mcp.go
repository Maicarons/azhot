package mcp

import (
	"api/all"
	"api/app"
	"api/config"
	"api/service"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// MCP协议基于JSON-RPC 2.0，用于AI模型与外部工具交互

// Request 表示MCP请求
type Request struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      string      `json:"id,omitempty"`
	Version string      `json:"jsonrpc"`
}

// Response 表示MCP响应
type Response struct {
	ID       string      `json:"id,omitempty"`
	Result   interface{} `json:"result,omitempty"`
	Error    *Error      `json:"error,omitempty"`
	Version  string      `json:"jsonrpc"`
}

// Error 表示MCP错误
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Tool 定义MCP工具
type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema Schema `json:"inputSchema"`
}

// Schema 定义工具输入模式
type Schema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required,omitempty"`
}

// ListPromptsResponse 响应结构
type ListPromptsResponse struct {
	Prompts []Prompt `json:"prompts"`
}

// Prompt 定义提示
type Prompt struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MCPHandler 处理MCP请求
type MCPHandler struct {
	service *service.HotSearchService
	config  *config.Config
	tools   map[string]Tool
}

// NewMCPHandler 创建新的MCP处理器
func NewMCPHandler(service *service.HotSearchService, config *config.Config) *MCPHandler {
	handler := &MCPHandler{
		service: service,
		config:  config,
		tools:   make(map[string]Tool),
	}
	
	// 注册可用的工具
	handler.registerTools()
	
	return handler
}

// registerTools 注册可用的工具
func (m *MCPHandler) registerTools() {
	// 注册获取热搜数据的工具
	m.tools["get_hot_search"] = Tool{
		Name:        "get_hot_search",
		Description: "获取各大平台的热搜数据，支持的平台包括baidu, bilibili, zhihu, weibo等",
		InputSchema: Schema{
			Type: "object",
			Properties: map[string]interface{}{
				"platform": map[string]interface{}{
					"type":        "string",
					"description": "平台名称，如baidu, bilibili, zhihu, weibo等",
				},
			},
			Required: []string{"platform"},
		},
	}
	
	// 注册获取所有平台热搜的工具
	m.tools["get_all_hot_search"] = Tool{
		Name:        "get_all_hot_search",
		Description: "获取所有平台的热搜数据聚合",
		InputSchema: Schema{
			Type:       "object",
			Properties: map[string]interface{}{},
		},
	}
	
	// 注册获取历史热搜数据的工具
	m.tools["get_history_data"] = Tool{
		Name:        "get_history_data",
		Description: "获取指定平台的历史热搜数据",
		InputSchema: Schema{
			Type: "object",
			Properties: map[string]interface{}{
				"platform": map[string]interface{}{
					"type":        "string",
					"description": "平台名称",
				},
				"date": map[string]interface{}{
					"type":        "string",
					"description": "日期，格式为YYYY-MM-DD",
				},
				"hour": map[string]interface{}{
					"type":        "string",
					"description": "小时，格式为HH",
				},
			},
			Required: []string{"platform", "date"},
		},
	}
}

// HandleRequest 处理MCP请求
func (m *MCPHandler) HandleRequest(requestBytes []byte) ([]byte, error) {
	var req Request
	if err := json.Unmarshal(requestBytes, &req); err != nil {
		return m.createErrorResponse("", -32700, "Parse error: unable to parse JSON")
	}

	switch req.Method {
	case "tools/list":
		return m.handleListTools(req.ID)
	case "tool/execute":
		return m.handleToolExecute(req)
	case "prompts/list":
		return m.handleListPrompts(req.ID)
	case "ping":
		return m.handlePing(req.ID)
	default:
		return m.createErrorResponse(req.ID, -32601, "Method not found")
	}
}

// handleListTools 处理工具列表请求
func (m *MCPHandler) handleListTools(id string) ([]byte, error) {
	tools := make([]Tool, 0, len(m.tools))
	for _, tool := range m.tools {
		tools = append(tools, tool)
	}

	response := Response{
		ID:      id,
		Result:  map[string]interface{}{"tools": tools},
		Version: "2.0",
	}

	return json.Marshal(response)
}

// handleToolExecute 处理工具执行请求
func (m *MCPHandler) handleToolExecute(req Request) ([]byte, error) {
	params, ok := req.Params.(map[string]interface{})
	if !ok {
		return m.createErrorResponse(req.ID, -32602, "Invalid params")
	}

	toolName, ok := params["name"].(string)
	if !ok {
		return m.createErrorResponse(req.ID, -32602, "Missing tool name")
	}

	_, exists := m.tools[toolName]
	if !exists {
		return m.createErrorResponse(req.ID, -32601, "Tool not found: "+toolName)
	}

	// 根据工具名称执行相应的操作
	switch toolName {
	case "get_hot_search":
		platform, ok := params["arguments"].(map[string]interface{})["platform"].(string)
		if !ok {
			return m.createErrorResponse(req.ID, -32602, "Missing platform argument")
		}
		return m.executeGetHotSearch(req.ID, platform)
	case "get_all_hot_search":
		return m.executeGetAllHotSearch(req.ID)
	case "get_history_data":
		args, ok := params["arguments"].(map[string]interface{})
		if !ok {
			return m.createErrorResponse(req.ID, -32602, "Missing arguments")
		}
		platform, platformOk := args["platform"].(string)
		date, dateOk := args["date"].(string)
		if !platformOk || !dateOk {
			return m.createErrorResponse(req.ID, -32602, "Missing required arguments: platform and date")
		}
		hour, _ := args["hour"].(string) // hour is optional
		return m.executeGetHistoryData(req.ID, platform, date, hour)
	default:
		return m.createErrorResponse(req.ID, -32601, "Unknown tool: "+toolName)
	}
}

// executeGetHotSearch 执行获取热搜数据的工具
func (m *MCPHandler) executeGetHotSearch(id, platform string) ([]byte, error) {
	// 导入app包来调用实际的API函数
	var result interface{}
	var err error

	// 使用反射或映射来调用相应的函数
	switch platform {
	case "baidu":
		result, err = callAppFunction("Baidu")
	case "bilibili":
		result, err = callAppFunction("Bilibili")
	case "zhihu":
		result, err = callAppFunction("Zhihu")
	case "weibo":
		result, err = callAppFunction("WeiboHot")
	case "360search":
		result, err = callAppFunction("Search360")
	case "acfun":
		result, err = callAppFunction("Acfun")
	case "csdn":
		result, err = callAppFunction("CSDN")
	case "dongqiudi":
		result, err = callAppFunction("Dongqiudi")
	case "douban":
		result, err = callAppFunction("Douban")
	case "douyin":
		result, err = callAppFunction("Douyin")
	case "github":
		result, err = callAppFunction("Github")
	case "guojiadili":
		result, err = callAppFunction("GuoJiadili")
	case "hupu":
		result, err = callAppFunction("Hupu")
	case "ithome":
		result, err = callAppFunction("Ithome")
	case "lishipin":
		result, err = callAppFunction("Lishipin")
	case "pengpai":
		result, err = callAppFunction("Pengpai")
	case "qqnews":
		result, err = callAppFunction("Qqnews")
	case "sougou":
		result, err = callAppFunction("Sougou")
	case "souhu":
		result, err = callAppFunction("Souhu")
	case "toutiao":
		result, err = callAppFunction("Toutiao")
	case "v2ex":
		result, err = callAppFunction("V2ex")
	case "wangyinews":
		result, err = callAppFunction("WangyiNews")
	case "xinjingbao":
		result, err = callAppFunction("Xinjingbao")
	case "renmin":
		result, err = callAppFunction("Renminwang")
	case "nanfang":
		result, err = callAppFunction("Nanfangzhoumo")
	case "360doc":
		result, err = callAppFunction("Doc360")
	case "cctv":
		result, err = callAppFunction("CCTV")
	case "quark":
		result, err = callAppFunction("Quark")
	default:
		return m.createErrorResponse(id, -32602, "Unsupported platform: "+platform)
	}

	if err != nil {
		return m.createErrorResponse(id, -32603, "Error calling API: "+err.Error())
	}

	response := Response{
		ID:      id,
		Result:  result,
		Version: "2.0",
	}

	return json.Marshal(response)
}

// executeGetAllHotSearch 执行获取所有平台热搜的工具
func (m *MCPHandler) executeGetAllHotSearch(id string) ([]byte, error) {
	result := all.All()

	response := Response{
		ID:      id,
		Result:  result,
		Version: "2.0",
	}

	return json.Marshal(response)
}

// callAppFunction 通过反射调用app包中的函数
func callAppFunction(funcName string) (interface{}, error) {
	// 获取app包中函数的反射值
	fn := reflect.ValueOf(getAppFunctionByName(funcName))
	if !fn.IsValid() || fn.Kind() != reflect.Func {
		return nil, fmt.Errorf("function %s not found", funcName)
	}

	// 调用函数
	results := fn.Call(nil)
	if len(results) != 2 { // 期望返回值, error
		return nil, fmt.Errorf("function %s does not return (interface{}, error)", funcName)
	}

	// 检查第二个返回值（error）
	err := results[1].Interface()
	if err != nil {
		if e, ok := err.(error); ok {
			return nil, e
		}
	}

	// 返回第一个返回值
	return results[0].Interface(), nil
}

// getAppFunctionByName 获取app包中指定名称的函数
func getAppFunctionByName(name string) interface{} {
	switch name {
	case "Baidu":
		return app.Baidu
	case "Bilibili":
		return app.Bilibili
	case "Zhihu":
		return app.Zhihu
	case "WeiboHot":
		return app.WeiboHot
	case "Search360":
		return app.Search360
	case "Acfun":
		return app.Acfun
	case "CSDN":
		return app.CSDN
	case "Dongqiudi":
		return app.Dongqiudi
	case "Douban":
		return app.Douban
	case "Douyin":
		return app.Douyin
	case "Github":
		return app.Github
	case "GuoJiadili":
		return app.Guojiadili // 注意：函数名是GuoJiadili，但变量名是guojiadili
	case "Hupu":
		return app.Hupu
	case "Ithome":
		return app.Ithome
	case "Lishipin":
		return app.Lishipin
	case "Pengpai":
		return app.Pengpai
	case "Qqnews":
		return app.Qqnews
	case "Sougou":
		return app.Sougou
	case "Souhu":
		return app.Souhu
	case "Toutiao":
		return app.Toutiao
	case "V2ex":
		return app.V2ex
	case "WangyiNews":
		return app.WangyiNews
	case "Xinjingbao":
		return app.Xinjingbao
	case "Renminwang":
		return app.Renminwang
	case "Nanfangzhoumo":
		return app.Nanfangzhoumo
	case "Doc360":
		return app.Doc360
	case "CCTV":
		return app.CCTV
	case "Quark":
		return app.Quark
	default:
		return nil
	}
}

// executeGetHistoryData 执行获取历史数据的工具
func (m *MCPHandler) executeGetHistoryData(id, platform, date, hour string) ([]byte, error) {
	var result interface{}
	var err error
	
	if hour != "" {
		// 获取指定日期和小时的历史数据
		result, err = m.service.GetHistoricalDataForWS(platform, date, hour)
	} else {
		// 获取指定日期的所有小时数据
		result, err = m.service.GetHistoricalDataByDateForWS(platform, date)
	}
	
	if err != nil {
		return m.createErrorResponse(id, -32603, "Error getting historical data: "+err.Error())
	}

	response := Response{
		ID:      id,
		Result:  result,
		Version: "2.0",
	}

	return json.Marshal(response)
}

// handleListPrompts 处理提示列表请求
func (m *MCPHandler) handleListPrompts(id string) ([]byte, error) {
	prompts := []Prompt{
		{
			Name:        "analyze_hot_search_trends",
			Description: "分析当前热搜趋势，识别热门话题和用户兴趣",
		},
		{
			Name:        "compare_platform_topics",
			Description: "比较不同平台的热门话题，分析差异和共同点",
		},
	}

	response := ListPromptsResponse{
		Prompts: prompts,
	}

	result := map[string]interface{}{
		"prompts": response.Prompts,
	}

	resp := Response{
		ID:      id,
		Result:  result,
		Version: "2.0",
	}

	return json.Marshal(resp)
}

// handlePing 处理ping请求
func (m *MCPHandler) handlePing(id string) ([]byte, error) {
	response := Response{
		ID:      id,
		Result:  map[string]interface{}{"message": "pong"},
		Version: "2.0",
	}

	return json.Marshal(response)
}

// createErrorResponse 创建错误响应
func (m *MCPHandler) createErrorResponse(id string, code int, message string) ([]byte, error) {
	response := Response{
		ID:      id,
		Error:   &Error{Code: code, Message: message},
		Version: "2.0",
	}

	return json.Marshal(response)
}

// RunMCPServerSTDIO 运行MCP服务器通过STDIO
func (m *MCPHandler) RunMCPServerSTDIO() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		
		// 处理输入的JSON-RPC请求
		response, err := m.HandleRequest([]byte(line))
		if err != nil {
			log.Printf("Error handling request: %v", err)
			continue
		}
		
		// 将响应输出到STDOUT
		fmt.Println(string(response))
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading stdin: %v", err)
	}
}

// RunMCPServerHTTP 运行MCP服务器通过HTTP
func (m *MCPHandler) RunMCPServerHTTP(port string) error {
	// 创建一个独立的Fiber应用作为MCP服务器
	app := fiber.New(fiber.Config{
		AppName:      "azhot MCP Server",
		ServerHeader: "azhot-mcp",
	})

	// 添加基本路由
	app.Post("/", func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		requestBytes, _ := json.Marshal(req)
		response, err := m.HandleRequest(requestBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Send(response)
	})

	log.Printf("Starting MCP HTTP server on port %s", port)
	return app.Listen(":" + port)
}