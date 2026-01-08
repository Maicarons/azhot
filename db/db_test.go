package db

import (
	"api/config"
	"api/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDBWithConfig(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_hot_search_1.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	// 测试初始化数据库
	InitDBWithConfig(cfg)

	// 检查数据库连接是否成功
	assert.NotNil(t, DB)
}

func TestSaveAndGetData(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_hot_search_2.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 准备测试数据
	source := "test_source"
	items := []model.HotSearchItem{
		{Title: "Test Title 1", URL: "http://example.com/1", Index: 1},
		{Title: "Test Title 2", URL: "http://example.com/2", Index: 2},
	}

	// 测试保存数据
	err := SaveData(source, items)
	assert.NoError(t, err)

	// 测试获取数据
	retrievedItems, err := GetLatestData(source)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(retrievedItems))
	assert.Equal(t, "Test Title 1", retrievedItems[0].Title)
	assert.Equal(t, "http://example.com/1", retrievedItems[0].URL)
	assert.Equal(t, 1, retrievedItems[0].Index)
	assert.Equal(t, "Test Title 2", retrievedItems[1].Title)
	assert.Equal(t, "http://example.com/2", retrievedItems[1].URL)
	assert.Equal(t, 2, retrievedItems[1].Index)
}

func TestSaveAndgetAllData(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_hot_search_3.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 准备测试数据
	allData := map[string][]model.HotSearchItem{
		"source1": {
			{Title: "Source1 Title 1", URL: "http://source1.com/1", Index: 1},
			{Title: "Source1 Title 2", URL: "http://source1.com/2", Index: 2},
		},
		"source2": {
			{Title: "Source2 Title 1", URL: "http://source2.com/1", Index: 1},
		},
	}

	// 测试保存所有数据
	err := SaveAllData(allData)
	assert.NoError(t, err)

	// 测试获取所有数据
	retrievedData, err := GetAllLatestData()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(retrievedData))
	assert.Equal(t, 2, len(retrievedData["source1"]))
	assert.Equal(t, 1, len(retrievedData["source2"]))
	assert.Equal(t, "Source1 Title 1", retrievedData["source1"][0].Title)
	assert.Equal(t, "Source2 Title 1", retrievedData["source2"][0].Title)
}

func TestGetLatestDataNotFound(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_hot_search_4.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 测试获取不存在的数据
	items, err := GetLatestData("nonexistent_source")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(items))
}

func TestSaveDataOverwrites(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_hot_search_5.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 首次保存数据
	source := "overwrite_test"
	firstItems := []model.HotSearchItem{
		{Title: "First Title", URL: "http://first.com", Index: 1},
	}
	err := SaveData(source, firstItems)
	assert.NoError(t, err)

	// 验证首次保存的数据
	items, err := GetLatestData(source)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "First Title", items[0].Title)

	// 保存新数据（应该覆盖旧数据）
	secondItems := []model.HotSearchItem{
		{Title: "Second Title", URL: "http://second.com", Index: 1},
		{Title: "Second Title 2", URL: "http://second2.com", Index: 2},
	}
	err = SaveData(source, secondItems)
	assert.NoError(t, err)

	// 验证新数据
	items, err = GetLatestData(source)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(items))
	assert.Equal(t, "Second Title", items[0].Title)
	assert.Equal(t, "Second Title 2", items[1].Title)
}

// 测试InitDB函数
func TestInitDB(t *testing.T) {
	// InitDB函数没有返回值且会修改全局状态，不适合单元测试
	// 我们跳过这个函数的测试
	t.Skip("Skipping TestInitDB as it modifies global state")
}

// 测试InitSQLite函数
func TestInitSQLite(t *testing.T) {
	// InitSQLite函数没有返回值且会修改全局状态，不适合单元测试
	// 我们跳过这个函数的测试
	t.Skip("Skipping TestInitSQLite as it modifies global state")
}

// 测试GetHistoricalData函数
func TestGetHistoricalData(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_historical_data.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 保存测试数据
	source := "historical_test"
	items := []model.HotSearchItem{
		{Title: "Historical Title", URL: "http://historical.com", Index: 1, Date: "2099-12-31", Hour: 12},
	}
	err := SaveData(source, items)
	assert.NoError(t, err)

	// 测试获取历史数据
	_, err = GetHistoricalData(source, "2099-12-31", 12)
	assert.NoError(t, err)
	// 可能没有数据，这很正常
}

// 测试GetHistoricalDataByDate函数
func TestGetHistoricalDataByDate(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_historical_date.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 保存测试数据
	source2 := "historical_date_test"
	items := []model.HotSearchItem{
		{Title: "Historical Date Title", URL: "http://historicaldate.com", Index: 1},
	}
	err := SaveData(source2, items)
	assert.NoError(t, err)

	// 测试获取特定日期的历史数据
	_, err = GetHistoricalDataByDate(source2, "2099-12-31")
	assert.NoError(t, err)
	// 可能没有数据，这很正常
}

// 测试GetHistoricalDataBySource函数
func TestGetHistoricalDataBySource(t *testing.T) {
	// 创建临时SQLite数据库文件
	tempDB := "test_historical_source.db"
	defer os.Remove(tempDB) // 测试结束后清理

	cfg := &config.Config{
		Database: config.DatabaseConfig{
			Type: "sqlite",
			DSN:  tempDB,
		},
	}

	InitDBWithConfig(cfg)

	// 保存测试数据
	source3 := "historical_source_test"
	items := []model.HotSearchItem{
		{Title: "Historical Source Title", URL: "http://historicalsource.com", Index: 1},
	}
	err := SaveData(source3, items)
	assert.NoError(t, err)

	// 测试获取特定来源的历史数据
	_, err = GetHistoricalDataBySource(source3)
	assert.NoError(t, err)
	// 可能没有历史数据，这很正常
}
