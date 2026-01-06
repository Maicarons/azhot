package db

import (
	"api/config"
	"api/model"
	"os"
	"time"

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
		InitSQLiteWithDSN(cfg.Database.DSN)
	}
}

// InitSQLite 初始化SQLite数据库
func InitSQLite() {

	DB = initSQLite("hot_search.db")

	log.Info("SQLite database initialized successfully")
}

// InitSQLiteWithDSN 使用DSN初始化SQLite数据库
func InitSQLiteWithDSN(dsn string) {

	DB = initSQLite(dsn)

	log.Info("SQLite database initialized successfully")
}

// initSQLite 初始化SQLite数据库的内部函数
func initSQLite(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + err.Error())
	}

	// 自动迁移模式
	err = db.AutoMigrate(&model.HotSearchItem{}, &model.HotSearchData{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
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
	// 为每个来源获取最新的数据批次
	data := make(map[string][]model.HotSearchItem)

	// 获取所有不同的来源
	var sources []string
	result := DB.Model(&model.HotSearchItem{}).Distinct("source").Pluck("source", &sources)
	if result.Error != nil {
		return nil, result.Error
	}

	// 对于每个来源，获取最新创建时间的数据
	for _, source := range sources {
		var latestCreatedAt time.Time
		// 获取此来源的最新创建时间
		result = DB.Model(&model.HotSearchItem{}).
			Where("source = ?", source).
			Order("created_at DESC").
			Limit(1).
			Pluck("created_at", &latestCreatedAt)
		if result.Error != nil {
			return nil, result.Error
		}

		// 获取此来源在最新时间的所有数据
		var items []model.HotSearchItem
		result = DB.Where("source = ? AND created_at = ?", source, latestCreatedAt).
			Order("item_index ASC").
			Find(&items)
		if result.Error != nil {
			return nil, result.Error
		}

		data[source] = items
	}

	return data, nil
}

// SaveData 保存数据到数据库
func SaveData(source string, items []model.HotSearchItem) error {
	// 删除旧数据
	if err := DB.Where("source = ?", source).Delete(&model.HotSearchItem{}).Error; err != nil {
		return err
	}

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
		// 删除旧数据
		if err := tx.Where("source = ?", source).Delete(&model.HotSearchItem{}).Error; err != nil {
			tx.Rollback()
			return err
		}

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
