package main

import (
	"net/http"

	"gopkg.in/pg.v5"
)

type application struct {
	cfg config
	db  *pg.DB
	srv *http.Server
}

type context struct {
	Title  string
	Static string
	Movies []Movie
	Genres []string
	Years  []string
	Count  int
}

func main() {
	app := new(application)
	app.getConfig()
	app.initDB()
	app.initServer()
}
