package db

import (
	"api/config"
	"api/model"
	"os"

	"github.com/gofiber/fiber/v2/log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接（保持向后兼容）
func InitDB() {
	// 检查是否使用MySQL环境变量
	dbType := os.Getenv("DB_TYPE") // sqlite 或 mysql
	if dbType == "mysql" {
		InitMySQL()
	} else {
		InitSQLite()
	}
}

// InitDBWithConfig 使用配置初始化数据库连接
func InitDBWithConfig(cfg *config.Config) {
	if cfg.Database.Type == "mysql" {
		InitMySQLWithConfig(cfg.Database.DSN)
	} else {
		InitSQLite()
	}
}

// InitSQLite 初始化SQLite数据库
func InitSQLite() {

	db, err := gorm.Open(sqlite.Open("hot_search.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + err.Error())
	}

	DB = db

	// 自动迁移模式
	err = DB.AutoMigrate(&model.HotSearchItem{}, &model.HotSearchData{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	log.Info("SQLite database initialized successfully")
}

// InitMySQL 初始化MySQL数据库（保持向后兼容）
func InitMySQL() {
	// 从环境变量获取MySQL连接参数
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/hot_search?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + err.Error())
	}

	DB = db

	// 自动迁移模式
	err = DB.AutoMigrate(&model.HotSearchItem{}, &model.HotSearchData{})
	if err != nil {
		log.Fatal("failed to migrate database: " + err.Error())
	}

	log.Info("MySQL database initialized successfully")
}

// InitMySQLWithConfig 使用配置初始化MySQL数据库
func InitMySQLWithConfig(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + err.Error())
	}

	DB = db

	// 自动迁移模式
	err = DB.AutoMigrate(&model.HotSearchItem{}, &model.HotSearchData{})
	if err != nil {
		log.Fatal("failed to migrate database: " + err.Error())
	}

	log.Info("MySQL database initialized successfully")
}

// GetLatestData 获取最新数据
func GetLatestData(source string) ([]model.HotSearchItem, error) {
	var items []model.HotSearchItem
	result := DB.Where("source = ?", source).Order("item_index ASC").Find(&items)
	return items, result.Error
}

// GetAllLatestData 获取所有最新数据
func GetAllLatestData() (map[string][]model.HotSearchItem, error) {
	var items []model.HotSearchItem
	result := DB.Order("source, item_index ASC").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	// 按来源分组
	data := make(map[string][]model.HotSearchItem)
	for _, item := range items {
		data[item.Source] = append(data[item.Source], item)
	}

	return data, nil
}

// SaveData 保存数据到数据库
func SaveData(source string, items []model.HotSearchItem) error {
	// 保存新数据
	for i := range items {
		items[i].Source = source
	}

	return DB.Create(&items).Error
}

// SaveAllData 保存所有数据
func SaveAllData(allData map[string][]model.HotSearchItem) error {
	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 逐个保存每个来源的数据
	for source, items := range allData {
		for i := range items {
			items[i].Source = source
		}
		if err := tx.Create(&items).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// GetHistoricalData 获取指定日期和小时的数据
func GetHistoricalData(source, date string, hour int) ([]model.HotSearchItem, error) {
	var items []model.HotSearchItem
	result := DB.Where("source = ? AND date = ? AND hour = ?", source, date, hour).Order("item_index ASC").Find(&items)
	return items, result.Error
}

// GetHistoricalDataByDate 获取指定日期的所有小时数据
func GetHistoricalDataByDate(source, date string) (map[int][]model.HotSearchItem, error) {
	var items []model.HotSearchItem
	result := DB.Where("source = ? AND date = ?", source, date).Order("hour, item_index ASC").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	// 按小时分组
	data := make(map[int][]model.HotSearchItem)
	for _, item := range items {
		data[item.Hour] = append(data[item.Hour], item)
	}

	return data, nil
}

// GetHistoricalDataBySource 获取指定来源的最新数据
func GetHistoricalDataBySource(source string) (map[string]map[int][]model.HotSearchItem, error) {
	var items []model.HotSearchItem
	result := DB.Where("source = ?", source).Order("date DESC, hour DESC, item_index ASC").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	// 按日期和小时分组
	data := make(map[string]map[int][]model.HotSearchItem)
	for _, item := range items {
		if data[item.Date] == nil {
			data[item.Date] = make(map[int][]model.HotSearchItem)
		}
		data[item.Date][item.Hour] = append(data[item.Date][item.Hour], item)
	}

	return data, nil
}
