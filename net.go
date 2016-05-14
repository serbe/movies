package main

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

func initRender() multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles("index", "templates/base.html", "templates/index.html")
	// r.AddFromFiles("react", "templates/react.html")

	return r
}

func (app *application) initServer() {
	r := gin.New()
	r.HTMLRender = initRender()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", root)

	r.GET("/movie", app.getOneMovieJSON)
	r.GET("/movies", app.getMoviesJSON)

	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.Run(":" + app.config.Web.Port)

	app.server = r
}
