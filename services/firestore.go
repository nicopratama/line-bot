package services

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func RegisterFirestore() *db.Client {
	ctx := context.Background()
	var ao map[string]interface{}
	conf := &firebase.Config{
		DatabaseURL:  "https://line-quiz.firebaseio.com",
		AuthOverride: &ao,
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

	ref := client.NewRef("/line-quiz")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)
	return client
}
