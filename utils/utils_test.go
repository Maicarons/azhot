package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMatches(t *testing.T) {
	t.Run("ExtractSimplePattern", func(t *testing.T) {
		text := "Hello world, hello golang"
		pattern := `hello`
		matches := ExtractMatches(text, pattern)

		// 由于正则表达式区分大小写，这里应该匹配到一个结果（"hello"）
		assert.Equal(t, 1, len(matches))
		assert.Equal(t, "hello", matches[0][0])

		// 测试区分大小写的匹配
		pattern = `Hello`
		matches = ExtractMatches(text, pattern)
		assert.Equal(t, 1, len(matches))
		assert.Equal(t, "Hello", matches[0][0])
	})

	t.Run("ExtractCaseInsensitivePattern", func(t *testing.T) {
		text := "Hello world, hello golang"
		pattern := `(?i)hello` // 不区分大小写的正则
		matches := ExtractMatches(text, pattern)

		assert.Equal(t, 2, len(matches))
		assert.Equal(t, "Hello", matches[0][0])
		assert.Equal(t, "hello", matches[1][0])
	})

	t.Run("ExtractWithGroups", func(t *testing.T) {
		text := "John is 25 years old, Jane is 30 years old"
		pattern := `(\w+) is (\d+) years old`
		matches := ExtractMatches(text, pattern)

		assert.Equal(t, 2, len(matches))
		assert.Equal(t, 3, len(matches[0])) // 整个匹配 + 2个组
		assert.Equal(t, "John is 25 years old", matches[0][0])
		assert.Equal(t, "John", matches[0][1])
		assert.Equal(t, "25", matches[0][2])
		assert.Equal(t, "Jane is 30 years old", matches[1][0])
		assert.Equal(t, "Jane", matches[1][1])
		assert.Equal(t, "30", matches[1][2])
	})

	t.Run("ExtractNoMatches", func(t *testing.T) {
		text := "Hello world"
		pattern := `xyz`
		matches := ExtractMatches(text, pattern)

		assert.Equal(t, 0, len(matches))
	})

	t.Run("ExtractEmptyText", func(t *testing.T) {
		text := ""
		pattern := `hello`
		matches := ExtractMatches(text, pattern)

		assert.Equal(t, 0, len(matches))
	})

	t.Run("ExtractEmptyPattern", func(t *testing.T) {
		text := "Hello world"
		pattern := ""
		matches := ExtractMatches(text, pattern)

		// 空模式会匹配每个位置，但通常会返回特殊结果
		assert.NotNil(t, matches)
	})

	t.Run("ExtractComplexPattern", func(t *testing.T) {
		text := "Contact us at info@example.com or support@test.org"
		pattern := `\b(\w+)@(\w+)\.(\w+)\b`
		matches := ExtractMatches(text, pattern)

		assert.Equal(t, 2, len(matches))
		assert.Equal(t, 4, len(matches[0])) // 整个匹配 + 3个组
		assert.Equal(t, "info@example.com", matches[0][0])
		assert.Equal(t, "info", matches[0][1])
		assert.Equal(t, "example", matches[0][2])
		assert.Equal(t, "com", matches[0][3])
		assert.Equal(t, "support@test.org", matches[1][0])
		assert.Equal(t, "support", matches[1][1])
		assert.Equal(t, "test", matches[1][2])
		assert.Equal(t, "org", matches[1][3])
	})
}
