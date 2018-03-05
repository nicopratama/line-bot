package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func RegisterFirestore() *firebase.App {
	opt := option.WithCredentialsFile("../config/line-quiz.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	return app
}
