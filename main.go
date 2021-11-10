package main

import "github.com/OkyWiliarso/news-api/database"

func main() {
	database.InitDB()

	defer database.DBCon.Close()
}
