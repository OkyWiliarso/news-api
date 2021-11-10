package database

import "github.com/jinzhu/gorm"

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	DBCon, err = gorm.Open("mysql", "root:qwerty123@tcp(localhost:3306)/news_db?parseTime=true")

	if err != nil {
		panic(err)
	}
}
