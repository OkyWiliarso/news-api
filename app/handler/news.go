package handler

import (
	"net/http"

	"github.com/OkyWiliarso/news-api/app/models"
	"github.com/jinzhu/gorm"
)

func GetAllnews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := []models.News{}
	db.Find(&news)
	respondJSON(w, http.StatusOK, news)
}
