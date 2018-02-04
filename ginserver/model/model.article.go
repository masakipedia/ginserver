// models.article.go

package model

import (
	"ginserver/ginserver/data"
	"ginserver/ginserver/idgenerator"

	"github.com/gin-gonic/gin"
)

// Return a list of all the articles
func GetAllArticles(c *gin.Context) []Article {
	dataArticles := data.GetAllArticles(c.Request)
	modelArticles := articles2model(dataArticles)
	return modelArticles
}

// Fetch an article based on the ID supplied
func GetArticleByID(c *gin.Context, id string) (Article, error) {
	dataArticle, err := data.GetArticle(c.Request, id)
	if err != nil {
		return Article{}, err
	}

	modelArticle := article2model(dataArticle)
	return modelArticle, nil
}

// Create a new article with the title and content provided
func CreateNewArticle(c *gin.Context, title, content string) (Article, error) {
	dataArticle := &data.Article{
		ID:      idgenerator.RandString6(6),
		Title:   title,
		Content: content,
	}

	if err := data.PutArticle(c.Request, dataArticle); err != nil {
		return Article{}, err
	}

	modelArticle := article2model(dataArticle)

	return modelArticle, nil
}
