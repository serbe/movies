package main

import (
	"github.com/labstack/echo"
	"gopkg.in/pg.v4"
)

type application struct {
	config   config
	server   *echo.Echo
	database *pg.DB
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
