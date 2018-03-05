package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func RegisterFirestore() *db.Client {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://databaseName.firebaseio.com",
	}
	opt := option.WithCredentialsFile("./config/line-quiz.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	return client
}
