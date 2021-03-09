package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// mapping all the popular css and js statics onto a static route
	static := r.Group("/static")
	static.Static("/jquery", "node_modules/jquery/dist")
	static.Static("/popper", "node_modules/popper.js/dist")
	static.Static("/bootswatch", "node_modules/bootswatch/dist")
	static.Static("/bootstrap-icons", "node_modules/bootstrap-icons/font")
	static.Static("/bootstrapjs", "node_modules/bootstrap/dist/js")
	static.Static("/bootstrap", "node_modules/bootstrap/dist/css")
	static.Static("/angular-route", "node_modules/angular-route")
	static.Static("/angular", "node_modules/angular")
	static.Static("/fortawesome", "node_modules/@fortawesome/fontawesome-free/css")
	static.Static("/fortawesomejs", "node_modules/@fortawesome/fontawesome-free/js")

	r.Run("0.0.0.0:8080")
}
