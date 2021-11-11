package main

import (
	"github.com/OkyWiliarso/news-api/app"
	"github.com/OkyWiliarso/news-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
