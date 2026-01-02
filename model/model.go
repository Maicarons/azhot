package model

import "time"

// HotSearchItem 表示单个热搜条目
type HotSearchItem struct {
	ID        uint      `json:"-" gorm:"primaryKey"`
	Source    string    `json:"source" gorm:"index"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Index     int       `json:"index" gorm:"column:item_index"`
	CreatedAt time.Time `json:"created_at"`
	Date      string    `json:"date" gorm:"index"` // 格式: YYYY-MM-DD
	Hour      int       `json:"hour" gorm:"index"` // 0-23
}

// HotSearchData 表示某个来源的完整热搜数据
type HotSearchData struct {
	ID        uint            `json:"-" gorm:"primaryKey"`
	Source    string          `json:"source" gorm:"index"`
	Items     []HotSearchItem `json:"items" gorm:"foreignKey:Source;references:Source"`
	CreatedAt time.Time       `json:"created_at"`
}
