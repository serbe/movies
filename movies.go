package main

import (
	"github.com/dinever/golf"
	"gopkg.in/pg.v4"
)

type application struct {
	config   config
	server   *golf.Application
	database *pg.DB
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
