package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", root)

	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.Run(":8080")
}
