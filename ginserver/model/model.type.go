package model

type Model struct {
	Articles []*Article
}

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
