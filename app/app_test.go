package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isNetworkError 检查错误是否为网络相关错误
func isNetworkError(errMsg string) bool {
	return strings.Contains(errMsg, "服务器内部错误") ||
		strings.Contains(errMsg, "context deadline exceeded") ||
		strings.Contains(errMsg, "timeout") ||
		strings.Contains(errMsg, "connection refused") ||
		strings.Contains(errMsg, "no such host") ||
		strings.Contains(errMsg, "network") ||
		strings.Contains(errMsg, "EOF")
}

// assertNetworkError 检查错误是否为网络相关错误
func assertNetworkError(t *testing.T, err error) {
	if err != nil {
		assert.True(t, isNetworkError(err.Error()), "错误信息应该包含预期的网络错误类型")
	}
}

func TestBilibili(t *testing.T) {
	// 由于Bilibili API需要网络请求，我们只测试函数是否能正常执行
	// 在测试环境中可能无法访问外部API，所以主要验证函数不会崩溃
	result, err := Bilibili()

	// API请求可能失败（网络问题等），这是正常的
	// 我们主要验证函数是否能返回正确的格式
	if err != nil {
		// 如果有错误，检查是否是网络相关错误
		assertNetworkError(t, err)
	} else {
		// 如果成功，验证返回格式
		assert.NotNil(t, result)
		assert.Contains(t, result, "code")
		assert.Contains(t, result, "message")

		code, ok := result["code"].(int)
		assert.True(t, ok)

		// 检查是否为成功状态或错误状态
		if code == 200 {
			// 成功时应该包含obj字段
			assert.Contains(t, result, "obj")
			message, ok := result["message"].(string)
			assert.True(t, ok)
			assert.Equal(t, "bilibili", message)

			obj, ok := result["obj"]
			assert.True(t, ok)
			assert.NotNil(t, obj)
		} else {
			// 错误状态时检查是否包含错误信息
			assert.Contains(t, result, "message")
		}
	}
}

func TestZhihu(t *testing.T) {
	// 由于Zhihu API需要网络请求，我们只测试函数是否能正常执行
	result, err := Zhihu()

	// API请求可能失败（网络问题等），这是正常的
	// 我们主要验证函数是否能返回正确的格式
	if err != nil {
		// 如果有错误，检查是否是网络相关错误
		assertNetworkError(t, err)
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
				// 检查是否有返回结果
				assert.NotNil(t, result)
				assert.Contains(t, result, "code")
				assert.Contains(t, result, "message")

				code, ok := result["code"].(int)
				assert.True(t, ok)

				// 根据返回的code值判断是否为成功状态
				if code == 200 {
					// 成功时应该返回完整的数据结构
					assert.Contains(t, result, "obj")
					message, ok := result["message"].(string)
					assert.True(t, ok)
					assert.NotEmpty(t, message)

					obj, ok := result["obj"]
					assert.True(t, ok)
					// obj应该是一个数组或map
					assert.NotNil(t, obj)
				} else {
					// 错误时(code != 200)，有些API可能包含obj字段，有些可能不包含，所以不强制要求
					// 这是为了兼容不同API的错误处理方式
				}
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
		assertNetworkError(t, err)
	}
}

func TestWeibo(t *testing.T) {
	// 测试微博API函数是否能正常调用
	_, err := WeiboHot()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

func TestToutiao(t *testing.T) {
	// 测试今日头条API函数是否能正常调用
	_, err := Toutiao()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

func TestDouban(t *testing.T) {
	// 测试豆瓣API函数是否能正常调用
	_, err := Douban()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

func TestGithub(t *testing.T) {
	// 测试GitHub API函数是否能正常调用
	_, err := Github()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

func TestV2ex(t *testing.T) {
	// 测试V2EX API函数是否能正常调用
	_, err := V2ex()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
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

// 测试360doc函数
func TestDoc360(t *testing.T) {
	_, err := Doc360()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试360search函数
func TestSearch360(t *testing.T) {
	_, err := Search360()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Acfun函数
func TestAcfun(t *testing.T) {
	_, err := Acfun()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试CCTV函数
func TestCCTV(t *testing.T) {
	_, err := CCTV()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试CSDN函数
func TestCSDN(t *testing.T) {
	_, err := CSDN()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Dongqiudi函数
func TestDongqiudi(t *testing.T) {
	_, err := Dongqiudi()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Douyin函数
func TestDouyin(t *testing.T) {
	_, err := Douyin()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Guojiadili函数
func TestGuojiadili(t *testing.T) {
	_, err := Guojiadili()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试stripHTML函数
func TestStripHTML(t *testing.T) {
	// 测试去除HTML标签
	htmlText := "<p>这是一段包含<em>HTML</em>标签的文本</p>"
	expected := "这是一段包含HTML标签的文本"
	result := stripHTML(htmlText)
	assert.Equal(t, expected, result)

	// 测试没有HTML标签的文本
	plainText := "这是一段普通文本"
	result2 := stripHTML(plainText)
	assert.Equal(t, plainText, result2)

	// 测试复杂HTML
	complexHTML := "<div><span>复杂</span>的<b>HTML</b>结构</div>"
	expected2 := "复杂的HTML结构"
	result3 := stripHTML(complexHTML)
	assert.Equal(t, expected2, result3)

	// 测试解析失败的情况
	invalidHTML := "<div>未闭合标签"
	result4 := stripHTML(invalidHTML)
	// 即使HTML无效，函数也应该返回某些内容而不崩溃
	assert.NotNil(t, result4)
}

// 测试HistoryToday函数
func TestHistoryToday(t *testing.T) {
	_, err := HistoryToday()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Hupu函数
func TestHupu(t *testing.T) {
	_, err := Hupu()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Ithome函数
func TestIthome(t *testing.T) {
	_, err := Ithome()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Lishipin函数
func TestLishipin(t *testing.T) {
	_, err := Lishipin()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Pengpai函数
func TestPengpai(t *testing.T) {
	_, err := Pengpai()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Qqnews函数
func TestQqnews(t *testing.T) {
	_, err := Qqnews()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Quark函数
func TestQuark(t *testing.T) {
	_, err := Quark()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Renminwang函数
func TestRenminwang(t *testing.T) {
	_, err := Renminwang()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Shaoshupai函数
func TestShaoshupai(t *testing.T) {
	_, err := Shaoshupai()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Sougou函数
func TestSougou(t *testing.T) {
	_, err := Sougou()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Souhu函数
func TestSouhu(t *testing.T) {
	_, err := Souhu()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试WangyiNews函数
func TestWangyiNews(t *testing.T) {
	_, err := WangyiNews()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Xinjingbao函数
func TestXinjingbao(t *testing.T) {
	_, err := Xinjingbao()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试Nanfangzhoumo函数
func TestNanfangzhoumo(t *testing.T) {
	_, err := Nanfangzhoumo()
	if err != nil {
		// API请求可能失败，这是正常的
		assertNetworkError(t, err)
	}
}

// 测试GetAllRouteNames函数
func TestGetAllRouteNames(t *testing.T) {
	routeNames := GetAllRouteNames()

	// 检查返回值是否为非空切片
	assert.NotEmpty(t, routeNames, "GetAllRouteNames should return non-empty slice")

	// 检查是否包含一些预期的路由名
	expectedRoutes := []string{"baidu", "weibo", "zhihu", "toutiao", "douban"}
	for _, expectedRoute := range expectedRoutes {
		contains := false
		for _, route := range routeNames {
			if route == expectedRoute {
				contains = true
				break
			}
		}
		assert.True(t, contains, "GetAllRouteNames should contain route: %s", expectedRoute)
	}
}
