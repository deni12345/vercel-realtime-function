package dao

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// init
type DAO struct {
	client *firestore.Client
}

func NewDAO(projectID string, credential []byte) *DAO {

	client, err := firestore.NewClient(context.Background(), "", option.WithCredentialsJSON(credential))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	return &DAO{client: client}
}
