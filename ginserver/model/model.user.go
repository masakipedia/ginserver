// models.user.go

package model

import (
	"errors"
	"ginserver/ginserver/data"
	"ginserver/ginserver/idgenerator"
	"strings"

	"github.com/gin-gonic/gin"
)

// Check if the username and password combination is valid
func IsUserValid(username, password string) bool {

	// TODO: User取得とう作る
	return true
}

// Register a new user with the given username and password
// NOTE: For this demo, we
func RegisterNewUser(c *gin.Context, username, password string) (User, error) {
	if message1, ok := IsPasswordAvailable(password); !ok {
		return User{}, errors.New(message1)
	}
	if message2, ok := IsUsernameAvailable(username); !ok {
		return User{}, errors.New(message2)
	}

	user := &data.User{
		ID:       idgenerator.RandString6(6),
		Username: username,
		Password: password,
	}
	if err := data.PutUser(c.Request, user); err != nil {
		return User{}, err
	}

	modelUser := user2model(user)

	return modelUser, nil
}

func IsPasswordAvailable(password string) (string, bool) {
	if strings.TrimSpace(password) == "" {
		return "The password can't be empty", false
	}
	return "", true
}

// Check if the supplied username is available
func IsUsernameAvailable(username string) (string, bool) {
	if strings.TrimSpace(username) == "" {
		return "The username can't be empty", false
	}
	return "", true
}
