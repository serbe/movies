package main

import (
	"github.com/gin-gonic/gin"
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
