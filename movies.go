package main

import (
	"gopkg.in/pg.v4"
	"net/http"
)

type application struct {
	config   config
	database *pg.DB
	server *http.Server
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
