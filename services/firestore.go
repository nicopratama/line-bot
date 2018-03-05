package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func RegisterFirestore() *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "line-quiz")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
