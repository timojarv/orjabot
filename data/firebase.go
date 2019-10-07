package data

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var fs *firestore.Client
var ctx = context.Background()

func init() {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SA")))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("error initializing app: %v", err)
	}

	fs, err = app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Firestore initialized")
}
