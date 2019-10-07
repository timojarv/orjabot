package data

import (
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"cloud.google.com/go/firestore"

	"google.golang.org/api/option"
)

var fs *firestore.Client

func init() {
	opt := option.WithCredentialsJson([]byte(os.Getenv("FIREBASE_SA")))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing app: %v", err)
	}

	fs, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Firestore initialized")
}
