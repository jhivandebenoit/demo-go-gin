package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {
	articles := getAllArticles()
	ctx.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"Title":   "Home Page",
			"payload": articles,
		},
	)

}

func getArticle(ctx *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the HTML method of the Context to render a template
			ctx.HTML(
				// Set the HTTP status to 200 (OK)
				http.StatusOK,
				// Use the index.html template
				"article.html",
				// Pass the data that the page uses
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// If the article is not found, abort with an error
			ctx.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}
