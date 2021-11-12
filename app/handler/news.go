package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OkyWiliarso/news-api/app/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := []models.News{}

	db.Preload("Tags").Find(&news)
	respondJSON(w, http.StatusOK, news)
}

func GetNewsByStatus(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := []models.News{}
	vars := mux.Vars(r)

	status := vars["status"]

	db.Where("status = ?", status).Find(&news)
	respondJSON(w, http.StatusOK, news)
}

func CreateNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	news := models.News{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&news); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Omit("Tags").Create(&news).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, news)
}

func UpdateNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	newsId, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	news := handleNewsNotFound(db, newsId, w, r)
	if news == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&news); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&news).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, news)
}

func DeleteNews(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	newsId, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	news := handleNewsNotFound(db, newsId, w, r)
	if news == nil {
		return
	}
	if err := db.Delete(&news).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func handleNewsNotFound(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.News {
	news := models.News{}
	if err := db.Preload("Tags").First(&news, models.News{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &news
}
