package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func RegisterFirestore() *firebase.App {
	opt := option.WithCredentialsFile("line-quiz-13f12b8df6c9.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	return app
}
