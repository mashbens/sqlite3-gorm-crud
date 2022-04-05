package main

import (
	"sqlite3-gorm/config"
	"sqlite3-gorm/route"
)

func main() {

	config.InitDB()

	app := route.New()

	app.Logger.Fatal(app.Start(":1322"))
}
