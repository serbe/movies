package main

import (
	"gopkg.in/pg.v4"
)

type application struct {
	config   config
	database *pg.DB
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
