package data

type (
	// Article is posted data.
	Article struct {
		ID      string `datastore:"-" goon:"id"`
		Title   string `datastore:"title,noindex"`
		Content string `datastore:"content,noindex"`
	}

	User struct {
		ID       string `datastore:"-" goon:"id"`
		Username string `datastore:"username,noindex"`
		Password string `datastore:"password,noindex"`
	}
)
