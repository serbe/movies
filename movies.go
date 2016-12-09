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

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
