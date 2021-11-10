package models

import "time"

type News struct {
	ID int `json:"id,primary_key"`

	Title   string `gorm:"size:255"`
	Content string `gorm:"size:255"`
	Status  string `gorm:"size:255"`

	Tags []Tags

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tags struct {
	ID int `json:"id,primary_key"` //this will be the primary key field

	NewsId int    `json:"news_id"`
	Topic  string `gorm:"size:255"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
