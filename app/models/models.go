package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type News struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	Tags    []Tags
}

type Tags struct {
	gorm.Model
	NewsId int    `json:"news_id"`
	Topic  string `json:"status"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&News{}, &Tags{})

	// db.Model(&News{})
	// db.Model(&Tags{}).AddForeignKey("news_id", "news(id)", "CASCADE", "RESTRICT")
	return db
}
