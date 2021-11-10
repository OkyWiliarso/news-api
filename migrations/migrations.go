package migrations

import (
	"github.com/OkyWiliarso/news-api/database"
	"github.com/OkyWiliarso/news-api/models"
)

func Migrate() {

	database.DBCon.AutoMigrate(models.News{}, models.Tags{})

	database.DBCon.Model(&models.News{})
	database.DBCon.Model(&models.Tags{}).AddForeignKey("news_id", "news(id)", "CASCADE", "RESTRICT")
}
