package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Renminwang 获取人民网热搜数据
//
//	@Summary		获取人民网热搜数据
//	@Description	获取人民网热门新闻排行榜
//	@Tags			renmin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/renmin [get]
func Renminwang() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// 更新URL到包含热点内容的页面
	// 根据提供的xpath，这些热点应该在主页或者指定的热点栏目页
	url := "http://www.people.com.cn/"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()
	
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}
	
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pageStr := string(pageBytes)
	
	// 根据提供的HTML结构，查找class="p6"的td元素
	re := regexp.MustCompile(`<td class="p6">((?s).*?)</td>`)
	matches := re.FindStringSubmatch(pageStr)
	
	if len(matches) < 2 {
		// 如果没找到class="p6"的td，尝试其他可能的热点区域
		// 查找可能包含热点新闻的li标签
		altPattern := `<li>\s*<a href="(.*?)"[^>]*?target="_blank"[^>]*?>([^<]*?)</a>`
		matched := utils.ExtractMatches(pageStr, altPattern)
		
		if len(matched) == 0 {
			return map[string]interface{}{
				"code":    500,
				"message": "未匹配到热点数据，可能页面结构已变更",
				"icon":    "http://www.people.com.cn/favicon.ico",
				"obj":     []map[string]interface{}{},
			}, nil
		}
		
		return extractResults(matched, "renmin")
	}
	
	// 提取td标签内的内容
	tdContent := matches[1]
	
	// 匹配<a>标签中的链接和标题
	// 注意：可能有多个链接在同一个<li>标签内
	linkPattern := `<a\s+href="(.*?)"[^>]*?target="_blank"[^>]*?>([^<]*?)</a>`
	matchedLinks := utils.ExtractMatches(tdContent, linkPattern)
	
	if len(matchedLinks) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未从热点区域内匹配到链接数据",
			"icon":    "http://www.people.com.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	
	return extractResults(matchedLinks, "renmin")
}

// extractResults 通用的结果提取函数
func extractResults(matched [][]string, source string) (map[string]interface{}, error) {
	var obj []map[string]interface{}
	seenTitles := make(map[string]bool) // 用于去重
	
	for index, item := range matched {
		if len(item) < 3 {
			continue
		}
		
		title := strings.TrimSpace(item[2])
		href := strings.TrimSpace(item[1])
		
		// 跳过空标题或重复标题
		if title == "" || seenTitles[title] {
			continue
		}
		
		seenTitles[title] = true
		
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = title
		result["url"] = normalizeURL(href)
		
		obj = append(obj, result)
	}
	
	// 如果没有有效数据
	if len(obj) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未提取到有效数据",
			"icon":    "http://www.people.com.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	
	api := map[string]interface{}{
		"code":    200,
		"message": source,
		"icon":    "http://www.people.com.cn/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api, nil
}

// normalizeURL 规范化URL
func normalizeURL(url string) string {
	if strings.HasPrefix(url, "//") {
		return "http:" + url
	}
	if strings.HasPrefix(url, "/") {
		return "http://www.people.com.cn" + url
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://www.people.com.cn" + url
	}
	return url
}