package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		// First, we add the headers with need to enable CORS
		// Make sure to adjust these headers to your needs
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		// c.Header("Cache-Control", "public, max-age=31536000")
		// c.Header("Content-Type", "application/json")
		// Second, we handle the OPTIONS problem
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			// Everytime we receive an OPTIONS request,
			// we just return an HTTP 200 Status Code
			// Like this, Angular can now do the real
			// request using any other method than OPTIONS
			c.AbortWithStatus(http.StatusOK)
		}
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// mapping all the popular css and js statics onto a static route
	// https://developers.google.com/web/fundamentals/performance/get-started/httpcaching-6
	// https://stackoverflow.com/questions/40714934/why-does-this-request-return-a-200from-cache-yet-others-return-304
	// https://github.com/gin-gonic/gin/issues/1222
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
	static.Static("/webfonts", "node_modules/@fortawesome/fontawesome-free/webfonts")

	r.Run("0.0.0.0:8080")
}
