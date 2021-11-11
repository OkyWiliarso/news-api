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
	a.Get("/news", a.GetAllnews)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) GetAllnews(w http.ResponseWriter, r *http.Request) {
	handler.GetAllnews(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
