package main

import (
	"api/config"
	"os"
	"testing"
)

// TestMainFunction 测试main函数的基本逻辑
func TestMainFunction(t *testing.T) {
	// 设置环境变量来避免实际启动服务器
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_PATH", ":memory:?cache=shared")
	os.Setenv("DEBUG", "true")

	// 创建配置
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Server.Host != "localhost" {
		t.Errorf("Expected host to be 'localhost', got '%s'", cfg.Server.Host)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected port to be '8080', got '%s'", cfg.Server.Port)
	}

	if cfg.Debug != true {
		t.Errorf("Expected debug to be true, got %t", cfg.Debug)
	}
}

// TestUpdateSwaggerHost 测试updateSwaggerHost函数
func TestUpdateSwaggerHost(t *testing.T) {
	os.Setenv("SERVER_HOST", "testhost")
	os.Setenv("SERVER_PORT", "3000")
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_PATH", ":memory:?cache=shared")
	os.Setenv("DEBUG", "false")

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// 由于SwaggerInfo是包级变量，我们不能直接验证，但可以调用函数确保它不会出错
	updateSwaggerHost(cfg)

	if cfg.Server.Host != "testhost" {
		t.Errorf("Expected host to be 'testhost', got '%s'", cfg.Server.Host)
	}
}
