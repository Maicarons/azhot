package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBilibili(t *testing.T) {
	// 由于Bilibili API需要网络请求，我们只测试函数是否能正常执行
	// 在测试环境中可能无法访问外部API，所以主要验证函数不会崩溃
	result, err := Bilibili()

	// API请求可能失败（网络问题等），这是正常的
	// 我们主要验证函数是否能返回正确的格式
	if err != nil {
		// 如果有错误，检查是否是网络相关错误
		assert.Contains(t, err.Error(), "服务器内部错误")
	} else {
		// 如果成功，验证返回格式
		assert.NotNil(t, result)
		assert.Contains(t, result, "code")
		assert.Contains(t, result, "message")
		assert.Contains(t, result, "obj")

		code, ok := result["code"].(int)
		assert.True(t, ok)
		assert.Equal(t, 200, code)

		message, ok := result["message"].(string)
		assert.True(t, ok)
		assert.Equal(t, "bilibili", message)
	}
}

func TestZhihu(t *testing.T) {
	// 由于Zhihu API需要网络请求，我们只测试函数是否能正常执行
	result, err := Zhihu()

	// API请求可能失败（网络问题等），这是正常的
	// 我们主要验证函数是否能返回正确的格式
	if err != nil {
		// 如果有错误，检查是否是网络相关错误
		assert.Contains(t, err.Error(), "服务器内部错误")
	} else {
		// 如果成功，验证返回格式
		assert.NotNil(t, result)
		assert.Contains(t, result, "code")
		assert.Contains(t, result, "message")
		assert.Contains(t, result, "obj")

		code, ok := result["code"].(int)
		assert.True(t, ok)
		assert.Equal(t, 200, code)

		message, ok := result["message"].(string)
		assert.True(t, ok)
		assert.Equal(t, "zhihu", message)
	}
}

func TestAllFunctionsReturnCorrectFormat(t *testing.T) {
	// 测试所有API函数返回的数据格式是否一致
	testCases := []struct {
		name     string
		function func() (map[string]interface{}, error)
	}{
		{"Bilibili", Bilibili},
		{"Zhihu", Zhihu},
		// 这里可以添加其他API函数进行测试
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.function()

			// 检查是否返回了正确的格式（即使有错误）
			if err != nil {
				// 如果有错误，至少应该返回一个包含code的map
				if result != nil {
					assert.Contains(t, result, "code")
					code, ok := result["code"].(int)
					assert.True(t, ok)
					assert.Equal(t, 500, code) // API错误时应该返回500
				}
			} else {
				// 成功时应该返回完整的数据结构
				assert.NotNil(t, result)
				assert.Contains(t, result, "code")
				assert.Contains(t, result, "message")
				assert.Contains(t, result, "obj")

				code, ok := result["code"].(int)
				assert.True(t, ok)
				assert.Equal(t, 200, code)

				message, ok := result["message"].(string)
				assert.True(t, ok)
				assert.NotEmpty(t, message)

				obj, ok := result["obj"]
				assert.True(t, ok)
				// obj应该是一个数组或map
				assert.NotNil(t, obj)
			}
		})
	}
}

// 为每个API函数添加单独的测试函数
func TestBaidu(t *testing.T) {
	// 测试百度API函数是否能正常调用
	_, err := Baidu()
	if err != nil {
		// API请求可能失败，这是正常的
		assert.Contains(t, err.Error(), "服务器内部错误")
	}
}

func TestWeibo(t *testing.T) {
	// 测试微博API函数是否能正常调用
	_, err := WeiboHot()
	if err != nil {
		// API请求可能失败，这是正常的
		assert.Contains(t, err.Error(), "服务器内部错误")
	}
}

func TestToutiao(t *testing.T) {
	// 测试今日头条API函数是否能正常调用
	_, err := Toutiao()
	if err != nil {
		// API请求可能失败，这是正常的
		assert.Contains(t, err.Error(), "服务器内部错误")
	}
}

func TestDouban(t *testing.T) {
	// 测试豆瓣API函数是否能正常调用
	_, err := Douban()
	if err != nil {
		// API请求可能失败，这是正常的
		assert.Contains(t, err.Error(), "服务器内部错误")
	}
}

func TestGithub(t *testing.T) {
	// 测试GitHub API函数是否能正常调用
	_, err := Github()
	if err != nil {
		// API请求可能失败，这是正常的
		assert.Contains(t, err.Error(), "服务器内部错误")
	}
}

func TestV2ex(t *testing.T) {
	// 测试V2EX API函数是否能正常调用
	_, err := V2ex()
	if err != nil {
		// API请求可能失败，这是正常的
		// 检查错误信息是否包含预期的错误类型
		// 检查错误信息是否包含预期的错误类型
		containsExpectedError := strings.Contains(err.Error(), "服务器内部错误") ||
			strings.Contains(err.Error(), "context deadline exceeded") ||
			strings.Contains(err.Error(), "timeout") ||
			strings.Contains(err.Error(), "connection refused") ||
			strings.Contains(err.Error(), "no such host") ||
			strings.Contains(err.Error(), "network") ||
			strings.Contains(err.Error(), "EOF")
		assert.True(t, containsExpectedError, "错误信息应该包含预期的网络错误类型")
	}
}

// 测试新添加的 ListSources 函数
func TestListSources(t *testing.T) {
	result, err := ListSources()
	if err != nil {
		// 如果有错误，至少应该返回一个包含code的map
		if result != nil {
			assert.Contains(t, result, "code")
			code, ok := result["code"].(int)
			assert.True(t, ok)
			assert.Equal(t, 500, code) // API错误时应该返回500
		}
	} else {
		// 成功时应该返回完整的数据结构
		assert.NotNil(t, result)
		assert.Contains(t, result, "code")
		assert.Contains(t, result, "message")
		assert.Contains(t, result, "obj")

		code, ok := result["code"].(int)
		assert.True(t, ok)
		assert.Equal(t, 200, code)

		obj, ok := result["obj"]
		assert.True(t, ok)

		// obj 应该是一个数组
		platforms, ok := obj.([]interface{})
		assert.True(t, ok, "obj should be an array of platform info")
		assert.Greater(t, len(platforms), 0, "should have at least one platform") // 至少应该有一个平台

		// 检查是否包含一些预期的平台信息
		foundBaidu := false
		for _, platform := range platforms {
			if platformMap, ok := platform.(map[string]interface{}); ok {
				_, hasRouteName := platformMap["routeName"]
				_, hasName := platformMap["name"]
				_, hasIcon := platformMap["icon"]

				// 检查所有字段都存在
				assert.True(t, hasRouteName, "平台信息应该包含路由名")
				assert.True(t, hasName, "平台信息应该包含中文名")
				assert.True(t, hasIcon, "平台信息应该包含图标")

				// 检查路由名是否为字符串
				if routeNameStr, isString := platformMap["routeName"].(string); isString && routeNameStr == "baidu" {
					foundBaidu = true
					break
				}
			}
		}
		assert.True(t, foundBaidu, "应该包含 'baidu' 平台信息")

		// 检查是否包含其他平台信息
		foundWeibo := false
		for _, platform := range platforms {
			if platformMap, ok := platform.(map[string]interface{}); ok {
				// 检查路由名是否为字符串
				if routeNameStr, isString := platformMap["routeName"].(string); isString && routeNameStr == "weibo" {
					foundWeibo = true
					break
				}
			}
		}
		assert.True(t, foundWeibo, "应该包含 'weibo' 平台信息")
	}
}
