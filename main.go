package main

import (
	"github.com/OkyWiliarso/news-api/database"
	"github.com/OkyWiliarso/news-api/migrations"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	database.InitDB()
	migrations.Migrate()

	defer database.DBCon.Close()
}
