package app

import (
	"api/utils"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

// Baidu 获取百度热搜数据
//
//	@Summary		获取百度热搜数据
//	@Description	获取百度热搜列表
//	@Tags			baidu
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/baidu [get]
func Baidu() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://top.baidu.com/board?tab=realtime"
	resp, err := client.Get(url)
	if err != nil {
		log.Error("http.Get error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("io.ReadAll error: %v", err)
		return nil, err
	}

	pattern := `<div\sclass="c-single-text-ellipsis">(.*?)</div?`
	// 注意：这里假设 utils.ExtractMatches 没有改变，如果它也返回 error，需要修改
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		// 添加边界检查
		if len(item) >= 2 {
			title := strings.TrimSpace(item[1])
			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": title,
				"url":   "https://www.baidu.com/s?wd=" + title,
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "baidu",
		"icon":    "https://www.baidu.com/favicon.ico", // 64 x 64
		"obj":     obj,
	}
	return api, nil
}
