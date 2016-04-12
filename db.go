package main

import (
	"gopkg.in/pg.v4"
)

func (app *application) initDB() {
	db := pg.Connect(&pg.Options{
		Database: app.config.Base.Dbname,
		User:     app.config.Base.User,
		Password: app.config.Base.Password,
		SSL:      app.config.Base.Sslmode,
	})
	app.database = db
}
