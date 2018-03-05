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
		DatabaseURL: "https://line-quiz.firebaseio.com",
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

	ref := client.NewRef("tes")
	type User struct {
		DateOfBirth string `json:"date_of_birth,omitempty"`
		FullName    string `json:"full_name,omitempty"`
		Nickname    string `json:"nickname,omitempty"`
	}

	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*User{
		"alanisawesome": &User{
			DateOfBirth: "June 23, 1912",
			FullName:    "Alan Turing",
		},
		"gracehop": &User{
			DateOfBirth: "December 9, 1906",
			FullName:    "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}

	return client
}
