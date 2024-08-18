package firebaseClient

import (
	"log"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func AddUser(user *User) bool {
	if client == nil {
		CreateClient()
	}

	_, err := client.Collection("Users").Doc(user.Email).Set(ctx, map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"gender":   user.Gender,
	})
	if err != nil {
		print(err, "\n")
		log.Fatal(err)
		return false
	}

	return true
}
