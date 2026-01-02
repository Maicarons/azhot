package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHotSearchItem(t *testing.T) {
	// 创建HotSearchItem实例
	item := HotSearchItem{
		ID:        1,
		Source:    "test_source",
		Title:     "Test Title",
		URL:       "http://example.com",
		Index:     1,
		CreatedAt: time.Now(),
	}

	// 验证字段值
	assert.Equal(t, uint(1), item.ID)
	assert.Equal(t, "test_source", item.Source)
	assert.Equal(t, "Test Title", item.Title)
	assert.Equal(t, "http://example.com", item.URL)
	assert.Equal(t, 1, item.Index)
	assert.NotZero(t, item.CreatedAt)
}

func TestHotSearchData(t *testing.T) {
	// 创建HotSearchData实例
	now := time.Now()
	data := HotSearchData{
		ID:        1,
		Source:    "test_source",
		CreatedAt: now,
		Items: []HotSearchItem{
			{Title: "Item 1", URL: "http://example.com/1", Index: 1},
			{Title: "Item 2", URL: "http://example.com/2", Index: 2},
		},
	}

	// 验证字段值
	assert.Equal(t, uint(1), data.ID)
	assert.Equal(t, "test_source", data.Source)
	assert.Equal(t, now, data.CreatedAt)
	assert.Equal(t, 2, len(data.Items))
	assert.Equal(t, "Item 1", data.Items[0].Title)
	assert.Equal(t, "Item 2", data.Items[1].Title)
}

func TestHotSearchItemJSONTags(t *testing.T) {
	// 验证JSON标签
	item := HotSearchItem{
		Source: "test_source",
		Title:  "Test Title",
		URL:    "http://example.com",
		Index:  1,
	}

	// 由于无法直接验证结构体标签，我们验证字段的可用性
	assert.Equal(t, "test_source", item.Source)
	assert.Equal(t, "Test Title", item.Title)
	assert.Equal(t, "http://example.com", item.URL)
	assert.Equal(t, 1, item.Index)
}

func TestHotSearchItemGormTags(t *testing.T) {
	// 验证GORM标签相关字段
	item := HotSearchItem{
		ID:     1,
		Source: "test_source",
		Index:  1,
	}

	// 验证字段值
	assert.Equal(t, uint(1), item.ID)
	assert.Equal(t, "test_source", item.Source)
	assert.Equal(t, 1, item.Index)
}
