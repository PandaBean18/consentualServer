package firebaseClient

import (
	"log"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
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
		//userId = *(userSnap.Data())
		*userId = userSnap.Data()["userId"].(string)
		return userId
	} else {
		_, err := client.Collection("Users").Doc(user.Email).Set(ctx, map[string]interface{}{
			"userId":   "someRandomIdForNow",
			"username": user.Username,
			"email":    user.Email,
			"gender":   user.Gender,
		})

		if err != nil {
			print(err, "\n")
			log.Fatal(err)
			return userId
		}
		*userId = "someRandomIdForNow"
	}

	return userId
}
