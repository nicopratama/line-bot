package services

import (
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
	"fmt"
)

func RegisterFirestore() *db.Client {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID:   "line-quiz",
		DatabaseURL: "https://line-quiz.firebaseio.com",
	}
	opt := option.WithCredentialsFile("./config/line-quiz-firebase-adminsdk-x7k1l-f532cb4e83.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	ref := client.NewRef("data")
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	fmt.Println(data)

	return client
}
