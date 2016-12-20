package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

type application struct {
	cfg config
	db  *sql.DB
	srv *http.Server
}

type context struct {
	Title  string
	Static string
	Movies []Movie
	Genres []string
	Years  []string
	Count  int64
}

func main() {
	app := new(application)
	app.getConfig()
	app.initDB()
	app.initServer()
}
