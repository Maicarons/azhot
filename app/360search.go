package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type search360Item struct {
	LongTitle string `json:"long_title"`
	Title     string `json:"title"`
	Score     string `json:"score"`
	Rank      string `json:"rank"`
}

// Search360 获取360搜索热搜数据
//
//	@Summary		获取360搜索热搜数据
//	@Description	获取360搜索热点排行榜
//	@Tags			search360
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/360search [get]
func Search360() (map[string]interface{}, error) {
	url := "https://ranks.hao.360.com/mbsug-api/hotnewsquery?type=news&realhot_limit=50"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Error("http.Get error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("io.ReadAll error: %v", err)
		return nil, err
	}

	var resultSlice []search360Item
	err = json.Unmarshal(pageBytes, &resultSlice)
	if err != nil {
		log.Error("json.Unmarshal error: %v", err)
		return nil, err
	}

	var obj []map[string]interface{}
	for _, item := range resultSlice {
		title := item.Title
		if item.LongTitle != "" {
			title = item.LongTitle
		}

		hot, err := strconv.ParseFloat(item.Score, 64)
		if err != nil {
			// 处理转换错误，可以给默认值或者跳过该项
			hot = 0
			// 或者使用日志记录错误但不中断程序
			// log.Printf("parse hot value error for item %s: %v", title, err)
		}

		// 将 hot/10000 格式化为一位小数的字符串，然后拼接 "万"
		hotValueStr := strconv.FormatFloat(hot/10000, 'f', 1, 64) + "万"
		obj = append(obj, map[string]interface{}{
			"index":    item.Rank,
			"title":    title,
			"hotValue": hotValueStr,
			"url":      "https://www.so.com/s?q=" + title,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "360search",
		"icon":    "https://ss.360tres.com/static/121a1737750aa53d.ico",
		"obj":     obj,
	}
	return api, nil
}
