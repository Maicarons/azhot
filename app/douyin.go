package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Douyinresponse struct {
	WordList []Douyindata `json:"word_list"`
}

type Douyindata struct {
	Title    string  `json:"word"`
	HotVaule float64 `json:"hot_value"`
}

// Douyin 获取抖音热搜数据
//
//	@Summary		获取抖音热搜数据
//	@Description	获取抖音热搜列表
//	@Tags			douyin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/douyin [get]
func Douyin() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	urlStr := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := client.Get(urlStr)
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

	var resultMap Douyinresponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		log.Error("json.Unmarshal error: %v", err)
		return nil, err
	}

	// 检查数据是否为空
	if len(resultMap.WordList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.WordList {
		// URL 编码标题，确保特殊字符正确处理
		encodedTitle := url.QueryEscape(item.Title)

		hotValue := ""
		if item.HotVaule > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.HotVaule/10000)
		}

		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      "https://www.douyin.com/search/" + encodedTitle,
			"hotValue": hotValue,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "douyin",
		"icon":    "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
