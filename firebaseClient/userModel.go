package firebaseClient

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Dob      string `json:"dob"`
	Age      string `json:"age"`
}

func createUserId() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	id := base64.StdEncoding.EncodeToString(bytes)

	return id[:(len(id) - 1)], nil
}

func AddUser(user *User, userId *string) *string {
	if client == nil {
		CreateClient()
	}

	userSnap, error := client.Collection("Users").Doc(user.Email).Get(ctx)

	if error != nil {
		log.Print(error)
	}

	if userSnap.Exists() {
		*userId = userSnap.Data()["userId"].(string)
		return userId
	} else {
		id, idErr := createUserId()
		if idErr != nil {
			log.Print("Error while creating userId.")
			return userId
		}

		print(user.Dob)
		_, err := client.Collection("Users").Doc(user.Email).Set(ctx, map[string]interface{}{
			"userId":   id,
			"username": user.Username,
			"email":    user.Email,
			"gender":   user.Gender,
			"dob":      user.Dob,
			"age":      user.Age,
		})

		if err != nil {
			print(err, "\n")
			log.Fatal(err)
			return userId
		}
		*userId = id
	}

	return userId
}
