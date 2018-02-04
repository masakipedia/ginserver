package model

import "ginserver/ginserver/data"

// Articles2models assigns the value of the matching structure of the field name.
func articles2model(dataArticles []*data.Article) []Article {
	modelArticles := make([]Article, len(dataArticles))
	for i, d := range dataArticles {
		modelArticles[i] = article2model(d)
	}
	return modelArticles
}

// Article2model assigns the value of the matching structure of the field name.
func article2model(dataArticles *data.Article) Article {
	return Article{
		Title:   dataArticles.Title,
		Content: dataArticles.Content,
	}
}

// Article2model assigns the value of the matching structure of the field name.
func user2model(dataUser *data.User) User {
	return User{
		ID:       dataUser.ID,
		Username: dataUser.Username,
		// Do not enter password.
	}
}
