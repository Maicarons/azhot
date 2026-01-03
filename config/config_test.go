package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// 测试默认配置
	t.Run("DefaultConfig", func(t *testing.T) {
		config, err := LoadConfig()
		assert.NoError(t, err)
		assert.Equal(t, "localhost", config.Server.Host)
		assert.Equal(t, "8080", config.Server.Port)
		assert.Equal(t, "sqlite", config.Database.Type)
		assert.Equal(t, "root:password@tcp(127.0.0.1:3306)/hot_search?charset=utf8mb4&parseTime=True&loc=Local", config.Database.DSN)
	})

	// 测试环境变量配置
	t.Run("EnvConfig", func(t *testing.T) {
		// 设置环境变量
		os.Setenv("SERVER_HOST", "testhost")
		os.Setenv("SERVER_PORT", "9090")
		os.Setenv("DB_TYPE", "mysql")
		os.Setenv("MYSQL_DSN", "test:test@tcp(localhost:3306)/testdb")

		config, err := LoadConfig()
		assert.NoError(t, err)
		assert.Equal(t, "testhost", config.Server.Host)
		assert.Equal(t, "9090", config.Server.Port)
		assert.Equal(t, "mysql", config.Database.Type)
		assert.Equal(t, "test:test@tcp(localhost:3306)/testdb", config.Database.DSN)
		assert.Equal(t, "", config.CORS.AllowOrigins) // 默认为空

		// 设置CORS环境变量
		os.Setenv("CORS_ALLOW_ORIGINS", "http://example.com,https://example.com")
		configWithCORS, err := LoadConfig()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com,https://example.com", configWithCORS.CORS.AllowOrigins)

		// 清理环境变量
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("DB_TYPE")
		os.Unsetenv("MYSQL_DSN")
		os.Unsetenv("CORS_ALLOW_ORIGINS")
	})
}

func TestGetServerAddress(t *testing.T) {
	config := &Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: "8080",
		},
	}

	address := config.GetServerAddress()
	assert.Equal(t, "localhost:8080", address)
}
