package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags    []Tags `gorm:"many2many:news_tags;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tags struct {
	ID    int    `json:"id"`
	Topic string `json:"topic"`
	News  []News `gorm:"many2many:news_tags;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&News{}, &Tags{})

	return db
}
