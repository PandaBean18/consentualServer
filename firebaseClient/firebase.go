package firebaseClient

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"github.com/joho/godotenv"
)

func CreateClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading the env file: ", err)
	}

	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatal("Error while initializing app: ", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("Error while initializing firestore: ", err)
	}

	defer client.Close()
}
