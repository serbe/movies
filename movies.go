package main

import (
	"github.com/labstack/echo"
	// "github.com/labstack/echo/engine/standard"
	"gopkg.in/pg.v4"
)

type application struct {
	config   config
	server   *gin.Engine
	database *pg.DB
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initServer()
}
