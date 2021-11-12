package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OkyWiliarso/news-api/app/handler"
	"github.com/OkyWiliarso/news-api/app/models"
	"github.com/OkyWiliarso/news-api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name)

	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	a.DB = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

}

func (a *App) setRouters() {
	a.Get("/news", a.GetAllNews)
	a.Post("/news", a.CreateNews)
	a.Put("/news/{id}", a.UpdateNews)
	a.Delete("/news/{id}", a.DeleteNews)

	a.Get("/tags", a.GetAllTags)
	a.Post("/tags", a.CreateTags)
	a.Put("/tags/{id}", a.UpdateTags)
	a.Delete("/tags/{id}", a.DeleteTags)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) GetAllNews(w http.ResponseWriter, r *http.Request) {
	handler.GetAllNews(a.DB, w, r)
}

func (a *App) CreateNews(w http.ResponseWriter, r *http.Request) {
	handler.CreateNews(a.DB, w, r)
}

func (a *App) UpdateNews(w http.ResponseWriter, r *http.Request) {
	handler.UpdateNews(a.DB, w, r)
}

func (a *App) DeleteNews(w http.ResponseWriter, r *http.Request) {
	handler.DeleteNews(a.DB, w, r)
}

func (a *App) GetAllTags(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTags(a.DB, w, r)
}

func (a *App) CreateTags(w http.ResponseWriter, r *http.Request) {
	handler.CreateTags(a.DB, w, r)
}

func (a *App) UpdateTags(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTags(a.DB, w, r)
}

func (a *App) DeleteTags(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTags(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
