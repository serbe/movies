package main

import (
	"html/template"
	
	"github.com/dinever/golf"	
	"gopkg.in/pg.v4"
)

type application struct {
	config   config
	server   *golf.Golf
	templates map[string]*template.Template
	database *pg.DB
}

func main() {
	app := application{}
	app.getConfig()
	app.initDB()
	app.initRender()
	app.initServer()
}
