package service

import (
	"api/config"
	"api/db"
	"api/model"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetFromDBOrFetch(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_service_hot_search_1.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 测试从API获取数据（数据库为空的情况）
	t.Run("FetchFromAPIWhenDBEmpty", func(t *testing.T) {
		result, err := service.GetFromDBOrFetch("微博") // 使用一个有效的源
		// 这里我们只测试函数是否能正常处理错误，因为实际API可能无法访问
		// 如果API无法访问，应该返回错误，否则应该返回数据
		if err != nil {
			// API请求可能失败，这在测试环境中是正常的
			assert.Contains(t, err.Error(), "服务器内部错误")
		} else {
			// 如果API成功，验证返回格式
			assert.NotNil(t, result)
			assert.Contains(t, result, "code")
		}
	})

	// 测试从数据库获取数据
	t.Run("FetchFromDB", func(t *testing.T) {
		// 首先手动保存一些测试数据到数据库
		testItems := []model.HotSearchItem{
			{Title: "Test Title", URL: "http://example.com", Index: 1},
		}
		err := db.SaveData("test_source", testItems)
		assert.NoError(t, err)

		// 从数据库获取数据
		result, err := service.GetFromDBOrFetch("test_source")
		assert.NoError(t, err)
		assert.Equal(t, 200, result["code"])
		assert.Equal(t, "test_source", result["message"])
		assert.NotNil(t, result["obj"])
	})
}

func TestGetAllFromDBOrFetch(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_service_hot_search_3.db" // 使用不同的文件名
	defer os.Remove(tempDB)                  // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 测试从数据库获取所有数据
	t.Run("FetchAllFromDB", func(t *testing.T) {
		// 首先手动保存一些测试数据到数据库
		testData := map[string][]model.HotSearchItem{
			"source1": {
				{Title: "Source1 Title", URL: "http://source1.com", Index: 1},
			},
			"source2": {
				{Title: "Source2 Title", URL: "http://source2.com", Index: 1},
			},
		}
		err := db.SaveAllData(testData)
		assert.NoError(t, err)

		// 从数据库获取所有数据
		result, err := service.GetAllFromDBOrFetch()
		assert.NoError(t, err)
		assert.Equal(t, 200, result["code"])
		assert.NotNil(t, result["obj"])
		obj, ok := result["obj"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, 2, len(obj))
	})
}

func TestConvertToHotSearchItems(t *testing.T) {
	service := &HotSearchService{}

	// 测试不同的API返回格式
	t.Run("ConvertFromSliceInterface", func(t *testing.T) {
		apiResult := map[string]interface{}{
			"obj": []interface{}{
				map[string]interface{}{"title": "Test Title", "url": "http://example.com", "index": 1.0},
			},
		}

		items := service.convertToHotSearchItems(apiResult)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, "Test Title", items[0].Title)
		assert.Equal(t, "http://example.com", items[0].URL)
		assert.Equal(t, 1, items[0].Index)
	})

	t.Run("ConvertFromSliceMap", func(t *testing.T) {
		apiResult := map[string]interface{}{
			"obj": []map[string]interface{}{
				{"title": "Test Title", "url": "http://example.com", "index": 1.0},
			},
		}

		items := service.convertToHotSearchItems(apiResult)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, "Test Title", items[0].Title)
		assert.Equal(t, "http://example.com", items[0].URL)
		assert.Equal(t, 1, items[0].Index)
	})

	t.Run("ConvertEmptyObj", func(t *testing.T) {
		apiResult := map[string]interface{}{
			"obj": []interface{}{},
		}

		items := service.convertToHotSearchItems(apiResult)
		assert.Equal(t, 0, len(items))
	})
}

func TestConvertFromDBItems(t *testing.T) {
	service := &HotSearchService{}

	// 测试将数据库项目转换为API返回格式
	t.Run("ConvertDBItemsToAPIFormat", func(t *testing.T) {
		dbItems := []model.HotSearchItem{
			{Title: "Test Title", URL: "http://example.com", Index: 1},
		}

		result := service.convertFromDBItems(dbItems, "test_source")
		assert.Equal(t, 200, result["code"])
		assert.Equal(t, "test_source", result["message"])
		assert.NotNil(t, result["obj"])

		obj, ok := result["obj"].([]map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, 1, len(obj))
		assert.Equal(t, "Test Title", obj[0]["title"])
		assert.Equal(t, "http://example.com", obj[0]["url"])
		// index 可能是 int 或 float64 类型，需要进行类型检查
		indexValue := obj[0]["index"]
		switch v := indexValue.(type) {
		case int:
			assert.Equal(t, 1, v)
		case float64:
			assert.Equal(t, 1, int(v))
		default:
			assert.Fail(t, "index should be int or float64")
		}
	})
}

func TestConvertFromAllDBItems(t *testing.T) {
	service := &HotSearchService{}

	// 测试将所有数据库项目转换为API返回格式
	t.Run("ConvertAllDBItemsToAPIFormat", func(t *testing.T) {
		dbData := map[string][]model.HotSearchItem{
			"source1": {
				{Title: "Source1 Title", URL: "http://source1.com", Index: 1},
			},
			"source2": {
				{Title: "Source2 Title", URL: "http://source2.com", Index: 1},
			},
		}

		result := service.convertFromAllDBItems(dbData)
		assert.Equal(t, 200, result["code"])
		assert.NotNil(t, result["obj"])

		obj, ok := result["obj"].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, 2, len(obj))

		source1Obj, ok := obj["source1"].([]map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, 1, len(source1Obj))
		assert.Equal(t, "Source1 Title", source1Obj[0]["title"])
	})
}

func TestConvertSourceDataToHotSearchItems(t *testing.T) {
	service := &HotSearchService{}

	// 测试从源数据转换为HotSearchItem
	t.Run("ConvertFromSliceInterface", func(t *testing.T) {
		sourceData := []interface{}{
			map[string]interface{}{"title": "Test Title", "url": "http://example.com", "index": 1.0},
		}

		items := service.convertSourceDataToHotSearchItems(sourceData)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, "Test Title", items[0].Title)
		assert.Equal(t, "http://example.com", items[0].URL)
		assert.Equal(t, 1, items[0].Index)
	})

	t.Run("ConvertFromSliceMap", func(t *testing.T) {
		sourceData := []map[string]interface{}{
			{"title": "Test Title", "url": "http://example.com", "index": 1.0},
		}

		items := service.convertSourceDataToHotSearchItems(sourceData)
		assert.Equal(t, 1, len(items))
		assert.Equal(t, "Test Title", items[0].Title)
		assert.Equal(t, "http://example.com", items[0].URL)
		assert.Equal(t, 1, items[0].Index)
	})

	t.Run("ConvertEmptyData", func(t *testing.T) {
		var sourceData []interface{}
		items := service.convertSourceDataToHotSearchItems(sourceData)
		assert.Equal(t, 0, len(items))
	})
}

func TestFetchDataFromAPI(t *testing.T) {
	service := &HotSearchService{}

	// 测试获取已知来源的数据
	t.Run("FetchKnownSource", func(t *testing.T) {
		result, err := service.FetchDataFromAPI("微博")
		// 这里我们只测试函数是否能处理请求，实际API可能无法访问
		// 如果API无法访问，应该返回错误，否则应该返回数据
		if err != nil {
			// API请求可能失败，这在测试环境中是正常的
			assert.Contains(t, err.Error(), "服务器内部错误")
		} else {
			// 如果API成功，验证返回格式
			assert.NotNil(t, result)
			assert.Contains(t, result, "code")
		}
	})

	// 测试获取未知来源的数据
	t.Run("FetchUnknownSource", func(t *testing.T) {
		result, err := service.FetchDataFromAPI("unknown_source")
		assert.NoError(t, err)
		assert.Equal(t, 200, result["code"])
		assert.Equal(t, "unknown_source", result["message"])
		assert.NotNil(t, result["obj"])
	})
}

// 测试fetchAPIData方法
func TestFetchAPIData(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_fetch_api_data.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 直接调用fetchAPIData方法，不期望出现panic或死锁
	// 由于这会尝试调用所有API，可能需要一些时间
	assert.NotPanics(t, func() {
		service.fetchAPIData()
	})
}

// 测试StartScheduler方法
func TestStartScheduler(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_scheduler.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 启动调度器，不期望出现panic
	assert.NotPanics(t, func() {
		service.StartScheduler()
	})
}

// 测试GetHistoricalDataHandler方法
func TestGetHistoricalDataHandler(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_handler.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 创建一个Fiber应用和上下文用于测试
	app := fiber.New()

	// 测试有效的参数
	app.Get("/test/:source/:date/:hour", func(c *fiber.Ctx) error {
		return service.GetHistoricalDataHandler(c)
	})

	// 发送请求
	req := httptest.NewRequest("GET", "/test/baidu/2023-01-01/12", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode) // 即使没有数据，路由处理程序也应该成功执行
}

// 测试GetHistoricalDataByDateHandler方法
func TestGetHistoricalDataByDateHandler(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_by_date_handler.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 创建一个Fiber应用和上下文用于测试
	app := fiber.New()

	// 测试有效的参数
	app.Get("/test/:source/:date", func(c *fiber.Ctx) error {
		return service.GetHistoricalDataByDateHandler(c)
	})

	// 发送请求
	req := httptest.NewRequest("GET", "/test/baidu/2023-01-01", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// 测试GetHistoricalDataBySourceHandler方法
func TestGetHistoricalDataBySourceHandler(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_by_source_handler.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 创建一个Fiber应用和上下文用于测试
	app := fiber.New()

	// 测试有效的参数
	app.Get("/test/:source", func(c *fiber.Ctx) error {
		return service.GetHistoricalDataBySourceHandler(c)
	})

	// 发送请求
	req := httptest.NewRequest("GET", "/test/baidu", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// 测试GetRouteNames方法
func TestGetRouteNames(t *testing.T) {
	service := &HotSearchService{}

	routeNames := service.GetRouteNames()

	// 检查返回值是否为非空切片
	assert.NotEmpty(t, routeNames, "GetRouteNames should return non-empty slice")

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
		assert.True(t, contains, "GetRouteNames should contain route: %s", expectedRoute)
	}
}

// 测试GetHistoricalDataForWS方法
func TestGetHistoricalDataForWS(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_for_ws.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 测试无效的小时参数
	result, err := service.GetHistoricalDataForWS("baidu", "2023-01-01", "invalid_hour")
	assert.NoError(t, err) // 函数应该处理错误而不是返回err
	assert.Equal(t, 500, result["code"])
	assert.Contains(t, result["message"], "小时参数格式错误")

	// 测试有效的参数（虽然可能没有数据）
	result2, err := service.GetHistoricalDataForWS("baidu", "2023-01-01", "12")
	assert.NoError(t, err)
	// 这里我们只检查函数是否正常执行，不检查返回的数据
	assert.NotNil(t, result2)
	assert.Contains(t, result2, "code")
}

// 测试GetHistoricalDataBySourceForWS方法
func TestGetHistoricalDataBySourceForWS(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_by_source_for_ws.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 测试有效的参数
	result, err := service.GetHistoricalDataBySourceForWS("baidu")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result, "code")
}

// 测试GetHistoricalDataByDateForWS方法
func TestGetHistoricalDataByDateForWS(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_history_by_date_for_ws.db"
	defer os.Remove(tempDB) // 测试结束后清理

	// 初始化数据库
	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}
	db.InitDBWithConfig(cfg)

	service := &HotSearchService{}

	// 测试有效的参数
	result, err := service.GetHistoricalDataByDateForWS("baidu", "2023-01-01")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, result, "code")
}
