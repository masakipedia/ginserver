package data

import (
	"net/http"

	"github.com/mjibson/goon"
)

func GetUser(r *http.Request, userID string) (*User, error) {
	g := goon.NewGoon(r)

	user := &User{
		ID: userID,
	}
	err := g.Get(user)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func PutUser(r *http.Request, user *User) error {
	g := goon.NewGoon(r)

	if _, err := g.Put(user); err != nil {
		return err
	}
	return nil
}
