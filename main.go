// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	// Start serving the application
	router.Run()

}

func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		ctx.JSON(http.StatusOK, data)
	case "application/html":
		ctx.HTML(http.StatusOK, templateName, data)
	case "application/xml":
		ctx.XML(http.StatusOK, data)

	default:
		ctx.HTML(http.StatusOK, templateName, data)

	}
}
