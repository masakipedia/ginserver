// handlers.article.go

package handler

import (
	"ginserver/ginserver/model"
	"ginserver/ginserver/render"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := model.GetAllArticles(c)

	// Call the render function with the name of the template to render
	render.Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render.Render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func GetArticle(c *gin.Context) {
	articleID := c.Param("article_id")

	// Check if the article exists
	if article, err := model.GetArticleByID(c, articleID); err == nil {
		// Call the render function with the title, article and the name of the
		// template
		render.Render(c, gin.H{
			"title":   article.Title,
			"payload": article}, "article.html")

	} else {
		// If the article is not found, abort with an error
		log.Printf("\n** data error **\n %v \n**\n\n", err)
		c.AbortWithError(http.StatusNotFound, err)
	}

}

func CreateArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := model.CreateNewArticle(c, title, content); err == nil {
		// If the article is created successfully, show success message
		render.Render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		log.Printf("\n** data error **\n %v \n**\n\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
