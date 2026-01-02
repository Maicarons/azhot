package service

import (
	"api/all"
	"api/app"
	"api/db"
	"api/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// HotSearchService 热搜服务
type HotSearchService struct{}

// GetFromDBOrFetch 从数据库获取最新数据，如果数据库为空则临时获取并保存
func (s *HotSearchService) GetFromDBOrFetch(source string) (map[string]interface{}, error) {
	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	// 首先尝试从数据库获取最新数据
	items, err := db.GetLatestData(dbSource)
	if err != nil {
		log.Errorf(fmt.Sprintf("从数据库获取 %s 数据失败: %v", source, err))
	}

	// 如果数据库中没有数据，则临时获取并保存
	if len(items) == 0 {
		log.Info("数据库中没有 " + source + " 数据，临时获取并保存...")
		result, err := s.FetchDataFromAPI(source)
		if err != nil {
			return nil, err
		}

		// 保存到数据库
		hotSearchItems := s.convertToHotSearchItems(result)
		if len(hotSearchItems) > 0 {
			err = db.SaveData(dbSource, hotSearchItems)
			if err != nil {
				log.Errorf(fmt.Sprintf("保存 %s 数据到数据库失败: %v", source, err))
			}
		}

		return result, nil
	}

	// 从数据库数据构建返回结果
	return s.convertFromDBItems(items, source), nil
}

// GetAllFromDBOrFetch 从数据库获取所有数据，如果数据库为空则临时获取并保存
func (s *HotSearchService) GetAllFromDBOrFetch() (map[string]interface{}, error) {
	// 首先尝试从数据库获取
	data, err := db.GetAllLatestData()
	if err != nil {
		log.Errorf("从数据库获取所有数据失败: %v", err)
	}

	// 如果数据库中没有数据，则临时获取并保存
	if len(data) == 0 {
		log.Info("数据库中没有数据，临时获取所有数据并保存...")
		result := all.All()

		// 转换数据并保存到数据库
		allData := make(map[string][]model.HotSearchItem)
		if obj, ok := result["obj"].(map[string]interface{}); ok {
			for source, sourceData := range obj {
				hotSearchItems := s.convertSourceDataToHotSearchItems(sourceData)
				// 使用路由名称作为源名称
				allData[source] = hotSearchItems
			}
		}

		if len(allData) > 0 {
			// 将路由名称转换为数据库源名称
			dbData := make(map[string][]model.HotSearchItem)
			for routeName, items := range allData {
				dbSource := s.convertRouteNameToDBSource(routeName)
				dbData[dbSource] = items
			}
			err = db.SaveAllData(dbData)
			if err != nil {
				log.Errorf(fmt.Sprintf("保存所有数据到数据库失败: %v", err))
			}
		}

		return result, nil
	}

	// 从数据库数据构建返回结果
	return s.convertFromAllDBItems(data), nil
}

// fetchAPIData 定时获取API数据并保存到数据库
func (s *HotSearchService) fetchAPIData() {
	log.Info("开始定时获取API数据...")

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	hour := currentTime.Hour()

	// 获取所有数据
	allResult := all.All()

	// 转换数据并保存到数据库
	allData := make(map[string][]model.HotSearchItem)
	if obj, ok := allResult["obj"].(map[string]interface{}); ok {
		for source, sourceData := range obj {
			hotSearchItems := s.convertSourceDataToHotSearchItems(sourceData)
			// 添加日期和小时信息
			for i := range hotSearchItems {
				hotSearchItems[i].Date = date
				hotSearchItems[i].Hour = hour
			}
			// 使用数据库源名称作为键
			dbSource := s.convertRouteNameToDBSource(source)
			allData[dbSource] = hotSearchItems
		}
	}

	if len(allData) > 0 {
		err := db.SaveAllData(allData)
		if err != nil {
			log.Errorf(fmt.Sprintf("定时保存所有数据到数据库失败: %v", err))
		} else {
			log.Info(fmt.Sprintf("定时获取API数据并保存到数据库完成，共保存 %d 个平台的数据，时间: %s %d:00", len(allData), date, hour))
		}
	}
}

// StartScheduler 启动定时任务
func (s *HotSearchService) StartScheduler() {
	// 立即执行一次
	s.fetchAPIData()

	// 每小时执行一次
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			s.fetchAPIData()
		}
	}()
}

// apiFunctionMap 将路由名称映射到API函数
var apiFunctionMap = map[string]func() (map[string]interface{}, error){
	"360search":  app.Search360,
	"bilibili":   app.Bilibili,
	"acfun":      app.Acfun,
	"csdn":       app.CSDN,
	"dongqiudi":  app.Dongqiudi,
	"douban":     app.Douban,
	"douyin":     app.Douyin,
	"github":     app.Github,
	"guojiadili": app.Guojiadili,
	"historytoday": app.HistoryToday,
	"hupu":       app.Hupu,
	"ithome":     app.Ithome,
	"lishipin":   app.Lishipin,
	"pengpai":    app.Pengpai,
	"qqnews":     app.Qqnews,
	"shaoshupai": app.Shaoshupai,
	"sougou":     app.Sougou,
	"toutiao":    app.Toutiao,
	"v2ex":       app.V2ex,
	"wangyinews": app.WangyiNews,
	"weibo":      app.WeiboHot,
	"xinjingbao": app.Xinjingbao,
	"zhihu":      app.Zhihu,
	"quark":      app.Quark,
	"kuake":      app.Quark,
	"souhu":      app.Souhu,
	"baidu":      app.Baidu,
	"renmin":     app.Renminwang,
	"nanfang":    app.Nanfangzhoumo,
	"360doc":     app.Doc360,
	"cctv":       app.CCTV,
}

// FetchDataFromAPI 根据来源获取API数据
func (s *HotSearchService) FetchDataFromAPI(source string) (map[string]interface{}, error) {
	// 检查映射中是否存在对应的函数
	if fn, exists := apiFunctionMap[source]; exists {
		return fn()
	}
	
	// 默认返回空结果
	return map[string]interface{}{
		"code":    200,
		"message": source,
		"obj":     []map[string]interface{}{},
	}, nil
}

// convertSourceDataToHotSearchItems 将源数据转换为HotSearchItem
func (s *HotSearchService) convertSourceDataToHotSearchItems(sourceData interface{}) []model.HotSearchItem {
	var hotSearchItems []model.HotSearchItem

	// 检查是否为 []interface{} 类型
	if items, ok := sourceData.([]interface{}); ok {
		for _, item := range items {
			if itemMap, ok := item.(map[string]interface{}); ok {
				title, _ := itemMap["title"].(string)
				url, _ := itemMap["url"].(string)
				index, _ := itemMap["index"].(float64)

				hotSearchItems = append(hotSearchItems, model.HotSearchItem{
					Title: title,
					URL:   url,
					Index: int(index),
				})
			}
		}
	} else if items, ok := sourceData.([]map[string]interface{}); ok {
		// 检查是否为 []map[string]interface{} 类型
		for _, item := range items {
			title, _ := item["title"].(string)
			url, _ := item["url"].(string)
			index, _ := item["index"].(float64)

			hotSearchItems = append(hotSearchItems, model.HotSearchItem{
				Title: title,
				URL:   url,
				Index: int(index),
			})
		}
	}

	return hotSearchItems
}

// convertToHotSearchItems 将API返回的数据转换为HotSearchItem
func (s *HotSearchService) convertToHotSearchItems(apiResult map[string]interface{}) []model.HotSearchItem {
	var items []model.HotSearchItem

	// 尝试不同类型的 obj 数据格式
	if obj, ok := apiResult["obj"].([]interface{}); ok {
		for _, item := range obj {
			if itemMap, ok := item.(map[string]interface{}); ok {
				title, _ := itemMap["title"].(string)
				url, _ := itemMap["url"].(string)
				index, _ := itemMap["index"].(float64)

				items = append(items, model.HotSearchItem{
					Title: title,
					URL:   url,
					Index: int(index),
				})
			}
		}
	} else if obj, ok := apiResult["obj"].([]map[string]interface{}); ok {
		for _, item := range obj {
			title, _ := item["title"].(string)
			url, _ := item["url"].(string)
			index, _ := item["index"].(float64)

			items = append(items, model.HotSearchItem{
				Title: title,
				URL:   url,
				Index: int(index),
			})
		}
	}

	return items
}

// convertFromDBItems 将数据库中的数据转换为API返回格式
func (s *HotSearchService) convertFromDBItems(items []model.HotSearchItem, source string) map[string]interface{} {
	var obj []map[string]interface{}

	for _, item := range items {
		obj = append(obj, map[string]interface{}{
			"index": item.Index,
			"title": item.Title,
			"url":   item.URL,
		})
	}

	return map[string]interface{}{
		"code":    200,
		"message": source,
		"obj":     obj,
	}
}

// convertFromAllDBItems 将所有数据库数据转换为API返回格式
func (s *HotSearchService) convertFromAllDBItems(data map[string][]model.HotSearchItem) map[string]interface{} {
	allObj := make(map[string]interface{})

	for source, items := range data {
		var obj []map[string]interface{}
		for _, item := range items {
			obj = append(obj, map[string]interface{}{
				"index": item.Index,
				"title": item.Title,
				"url":   item.URL,
			})
		}
		allObj[source] = obj
	}

	return map[string]interface{}{
		"code": 200,
		"obj":  allObj,
	}
}

// GetHistoricalDataHandler 获取指定日期和小时的历史数据
//
//	@Summary		获取指定日期和小时的历史数据
//	@Description	获取指定日期和小时的历史热搜数据
//	@Tags			HistoryAPI
//	@Accept			json
//	@Produce		json
//	@Param			source	path		string	true	"数据源名称"
//	@Param			date	path		string	true	"日期，格式：YYYY-MM-DD"
//	@Param			hour	path		int		true	"小时，0-23"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/history/{source}/{date}/{hour} [get]
func (s *HotSearchService) GetHistoricalDataHandler(c *fiber.Ctx) error {
	source := c.Params("source")
	date := c.Params("date")
	hourParam := c.Params("hour")
	hour, err := strconv.Atoi(hourParam)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "小时参数格式错误",
			"obj":     []map[string]interface{}{},
		})
	}

	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	items, err := db.GetHistoricalData(dbSource, date, hour)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     []map[string]interface{}{},
		})
	}

	if len(items) == 0 {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": fmt.Sprintf("未找到 %s 在 %s %d:00 的历史数据", source, date, hour),
			"obj":     []map[string]interface{}{},
		})
	}

	// 从数据库数据构建返回结果
	result := s.convertFromDBItems(items, source+"历史数据")
	return c.JSON(result)
}

// GetHistoricalDataByDateHandler 获取指定日期的所有小时数据
//
//	@Summary		获取指定日期的所有小时数据
//	@Description	获取指定日期的所有小时的历史热搜数据
//	@Tags			HistoryAPI
//	@Accept			json
//	@Produce		json
//	@Param			source	path		string	true	"数据源名称"
//	@Param			date	path		string	true	"日期，格式：YYYY-MM-DD"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/history/{source}/{date} [get]
func (s *HotSearchService) GetHistoricalDataByDateHandler(c *fiber.Ctx) error {
	source := c.Params("source")
	date := c.Params("date")

	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	data, err := db.GetHistoricalDataByDate(dbSource, date)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     map[string]interface{}{},
		})
	}

	// 转换数据格式
	result := make(map[string]interface{})
	for hour, items := range data {
		var obj []map[string]interface{}
		for _, item := range items {
			obj = append(obj, map[string]interface{}{
				"index": item.Index,
				"title": item.Title,
				"url":   item.URL,
			})
		}
		result[fmt.Sprintf("%02d:00", hour)] = obj
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"obj":  result,
	})
}

// GetHistoricalDataBySourceHandler 获取指定来源的最新历史数据
//
//	@Summary		获取指定来源的最新历史数据
//	@Description	获取指定来源的最新历史热搜数据
//	@Tags			HistoryAPI
//	@Accept			json
//	@Produce		json
//	@Param			source	path		string	true	"数据源名称"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Router			/history/{source} [get]
func (s *HotSearchService) GetHistoricalDataBySourceHandler(c *fiber.Ctx) error {
	source := c.Params("source")

	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	data, err := db.GetHistoricalDataBySource(dbSource)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     map[string]interface{}{},
		})
	}

	// 转换数据格式
	result := make(map[string]interface{})
	for date, hoursData := range data {
		hoursResult := make(map[string]interface{})
		for hour, items := range hoursData {
			var obj []map[string]interface{}
			for _, item := range items {
				obj = append(obj, map[string]interface{}{
					"index": item.Index,
					"title": item.Title,
					"url":   item.URL,
				})
			}
			hoursResult[fmt.Sprintf("%02d:00", hour)] = obj
		}
		result[date] = hoursResult
	}

	return c.JSON(fiber.Map{
		"code": 200,
		"obj":  result,
	})
}

// convertRouteNameToDBSource 将路由名称转换为数据库中存储的源名称
func (s *HotSearchService) convertRouteNameToDBSource(routeName string) string {
	// 现在直接返回路由名称，不再转换为中文名称
	return routeName
}

// GetRouteNames 获取所有可用的路由名称列表
func (s *HotSearchService) GetRouteNames() []string {
	return []string{
		"360search",
		"bilibili",
		"acfun",
		"csdn",
		"dongqiudi",
		"douban",
		"douyin",
		"github",
		"guojiadili",
		"history",
		"hupu",
		"ithome",
		"lishipin",
		"pengpai",
		"qqnews",
		"shaoshupai",
		"sougou",
		"toutiao",
		"v2ex",
		"wangyinews",
		"weibo",
		"xinjingbao",
		"zhihu",
		"quark",
		"souhu",
		"baidu",
		"renmin",
		"nanfang",
		"360doc",
		"cctv",
	}
}

// GetHistoricalDataForWS 获取指定日期和小时的历史数据用于WebSocket
func (s *HotSearchService) GetHistoricalDataForWS(source, date, hourParam string) (map[string]interface{}, error) {
	hour, err := strconv.Atoi(hourParam)
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": "小时参数格式错误",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	items, err := db.GetHistoricalData(dbSource, date, hour)
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     []map[string]interface{}{},
		}, nil
	}

	if len(items) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": fmt.Sprintf("未找到 %s 在 %s %d:00 的历史数据", source, date, hour),
			"obj":     []map[string]interface{}{},
		}, nil
	}

	// 从数据库数据构建返回结果
	result := s.convertFromDBItems(items, source+"历史数据")
	return result, nil
}

// GetHistoricalDataByDateForWS 获取指定日期的所有小时数据用于WebSocket
func (s *HotSearchService) GetHistoricalDataByDateForWS(source, date string) (map[string]interface{}, error) {
	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	data, err := db.GetHistoricalDataByDate(dbSource, date)
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     map[string]interface{}{},
		}, nil
	}

	// 转换数据格式
	result := make(map[string]interface{})
	for hour, items := range data {
		var obj []map[string]interface{}
		for _, item := range items {
			obj = append(obj, map[string]interface{}{
				"index": item.Index,
				"title": item.Title,
				"url":   item.URL,
			})
		}
		result[fmt.Sprintf("%02d:00", hour)] = obj
	}

	return map[string]interface{}{
		"code": 200,
		"obj":  result,
	}, nil
}

// GetHistoricalDataBySourceForWS 获取指定来源的最新历史数据用于WebSocket
func (s *HotSearchService) GetHistoricalDataBySourceForWS(source string) (map[string]interface{}, error) {
	// 将路由名称转换为数据库中存储的源名称
	dbSource := s.convertRouteNameToDBSource(source)

	data, err := db.GetHistoricalDataBySource(dbSource)
	if err != nil {
		return map[string]interface{}{
			"code":    500,
			"message": "服务器内部错误: " + err.Error(),
			"obj":     map[string]interface{}{},
		}, nil
	}

	// 转换数据格式
	result := make(map[string]interface{})
	for date, hoursData := range data {
		hoursResult := make(map[string]interface{})
		for hour, items := range hoursData {
			var obj []map[string]interface{}
			for _, item := range items {
				obj = append(obj, map[string]interface{}{
					"index": item.Index,
					"title": item.Title,
					"url":   item.URL,
				})
			}
			hoursResult[fmt.Sprintf("%02d:00", hour)] = obj
		}
		result[date] = hoursResult
	}

	return map[string]interface{}{
		"code": 200,
		"obj":  result,
	}, nil
}
