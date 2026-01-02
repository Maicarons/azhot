package mcp

import (
	"api/config"
	"api/service"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
)

// SetupMCPRoutes 设置MCP相关路由
func SetupMCPRoutes(app *fiber.App, service *service.HotSearchService, cfg *config.Config) {
	// 创建MCP处理器
	mcpHandler := NewMCPHandler(service, cfg)

	// MCP HTTP端点 - 用于调试和直接访问
	mcpGroup := app.Group("/mcp")

	// 工具列表端点
	mcpGroup.Get("/tools", func(c *fiber.Ctx) error {
		request := Request{
			Method:  "tools/list",
			ID:      "http-request",
			Version: "2.0",
		}
		requestBytes, _ := json.Marshal(request)
		response, err := mcpHandler.HandleRequest(requestBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Send(response)
	})

	// 工具执行端点
	mcpGroup.Post("/tool/execute", func(c *fiber.Ctx) error {
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request format",
			})
		}

		// 确保方法是tool/execute
		req.Method = "tool/execute"
		if req.Version == "" {
			req.Version = "2.0"
		}

		requestBytes, _ := json.Marshal(req)
		response, err := mcpHandler.HandleRequest(requestBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Send(response)
	})

	// 提示列表端点
	mcpGroup.Get("/prompts", func(c *fiber.Ctx) error {
		request := Request{
			Method:  "prompts/list",
			ID:      "http-request",
			Version: "2.0",
		}
		requestBytes, _ := json.Marshal(request)
		response, err := mcpHandler.HandleRequest(requestBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Send(response)
	})

	// Ping端点
	mcpGroup.Get("/ping", func(c *fiber.Ctx) error {
		request := Request{
			Method:  "ping",
			ID:      "http-request",
			Version: "2.0",
		}
		requestBytes, _ := json.Marshal(request)
		response, err := mcpHandler.HandleRequest(requestBytes)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Send(response)
	})

	// MCP发现端点 - 提供MCP服务器元数据
	mcpGroup.Get("/.well-known/mcp-info", func(c *fiber.Ctx) error {
		info := map[string]interface{}{
			"version": "1.0",
			"name":    "azhot MCP Server",
			"description": "MCP server for azhot - Hot Search API Aggregation Service",
			"tools": []string{"get_hot_search", "get_all_hot_search", "get_history_data"},
			"prompts": []string{"analyze_hot_search_trends", "compare_platform_topics"},
		}
		return c.JSON(info)
	})

	// 如果配置中启用了MCP STDIO服务器，则启动它
	if cfg.MCP != nil && cfg.MCP.STDIOEnabled {
		go func() {
			log.Println("Starting MCP STDIO server...")
			mcpHandler.RunMCPServerSTDIO()
		}()
	}
	
	// 如果配置中启用了MCP HTTP服务器，则启动它
	if cfg.MCP != nil && cfg.MCP.HTTPEnabled {
		go func() {
			log.Println("Starting MCP HTTP server on port", cfg.MCP.Port)
			err := mcpHandler.RunMCPServerHTTP(cfg.MCP.Port)
			if err != nil {
				log.Printf("Error starting MCP HTTP server: %v", err)
			}
		}()
	}
}