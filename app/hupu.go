package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Hupu 获取虎扑热搜数据
//
//	@Summary		获取虎扑热搜数据
//	@Description	获取虎扑热门体育资讯排行榜
//	@Tags			hupu
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/hupu [get]
func Hupu() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.hupu.com/"
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

	pattern := `<a\s+href="([^"]+)"[^>]+>\s*<div[^>]+>\s*<div[^>]+>\d+</div>\s*<div[^>]+>(.*?)</div>`
	matches := utils.ExtractMatches(string(pageBytes), pattern)

	// 检查是否匹配到数据
	if len(matches) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://www.hupu.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range matches {
		// 添加边界检查
		if len(item) >= 3 {
			url := item[1]
			title := item[2]

			// 确保 URL 是完整的
			if len(url) > 0 && url[0] == '/' {
				url = "https://www.hupu.com" + url
			}

			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": title,
				"url":   url,
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "hupu",
		"icon":    "https://www.hupu.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
