package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OkyWiliarso/news-api/app/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tags := []models.Tags{}
	db.Find(&tags)
	respondJSON(w, http.StatusOK, tags)
}

func CreateTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	tags := models.Tags{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&tags).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, tags)
}

func UpdateTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tagsId, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tags := handleTagsNotFound(db, tagsId, w, r)
	if tags == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tags); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&tags).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, tags)
}

func DeleteTags(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tagsId, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tags := handleTagsNotFound(db, tagsId, w, r)
	if tags == nil {
		return
	}
	if err := db.Delete(&tags).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func handleTagsNotFound(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.Tags {
	tags := models.Tags{}
	if err := db.First(&tags, models.Tags{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &tags
}
