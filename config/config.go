package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// CORSConfig CORS配置
type CORSConfig struct {
	AllowOrigins string // 允许的跨域请求来源
}

// Config 应用程序配置结构体
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	MCP      *MCPConfig
	CORS     CORSConfig
	Debug    bool
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host string
	Port string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Type string // "sqlite" 或 "mysql"
	DSN  string // 数据库连接字符串
}

// MCPConfig MCP服务器配置
type MCPConfig struct {
	STDIOEnabled bool   // 是否启用STDIO MCP服务器
	HTTPEnabled  bool   // 是否启用HTTP MCP服务器
	Port         string // HTTP MCP服务器端口
}

// LoadConfig 从环境变量或.env文件加载配置
func LoadConfig() (*Config, error) {
	// 加载 .env 文件（如果存在）
	godotenv.Load()

	config := &Config{
		Server: ServerConfig{
			Host: getEnvOrDefault("SERVER_HOST", "localhost"),
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Type: getEnvOrDefault("DB_TYPE", "sqlite"),
			DSN:  getEnvOrDefault("MYSQL_DSN", "root:password@tcp(127.0.0.1:3306)/hot_search?charset=utf8mb4&parseTime=True&loc=Local"),
		},
		CORS: CORSConfig{
			AllowOrigins: getEnvOrDefault("CORS_ALLOW_ORIGINS", ""),
		},
		MCP: &MCPConfig{
			STDIOEnabled: getEnvOrDefault("MCP_STDIO_ENABLED", "false") == "true",
			HTTPEnabled:  getEnvOrDefault("MCP_HTTP_ENABLED", "false") == "true",
			Port:         getEnvOrDefault("MCP_PORT", "8081"),
		},
		Debug: getEnvOrDefault("DEBUG", "false") == "true",
	}

	return config, nil
}

// getEnvOrDefault 获取环境变量，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetServerAddress 获取服务器完整地址
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}
