package data

import (
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
)

// GetAllArticles gets all articles from GAE.
func GetAllArticles(r *http.Request) []*Article {
	g := goon.NewGoon(r)
	q := datastore.NewQuery("Article")

	var articles []*Article
	g.GetAll(q, &articles)

	return articles
}

// GetArticle get a specific article.
func GetArticle(r *http.Request, articleID string) (*Article, error) {
	g := goon.NewGoon(r)

	article := &Article{
		ID: articleID,
	}
	err := g.Get(article)
	if err != nil {
		return &Article{}, err
	}
	return article, nil
}

// PutArticle posts an article.
func PutArticle(r *http.Request, article *Article) error {
	g := goon.NewGoon(r)

	if _, err := g.Put(article); err != nil {
		return err
	}
	return nil
}
