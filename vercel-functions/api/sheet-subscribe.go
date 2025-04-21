package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func getCredentials() ([]byte, error) {
	googleCredentials := os.Getenv("GOOGLE_CREDENTIALS")
	if googleCredentials == "" {
		return nil, fmt.Errorf("GOOGLE_CREDENTIALS environment variable is not set")
	}

	decodedCredentials, err := base64.RawStdEncoding.DecodeString(googleCredentials)
	if err != nil {
		return nil, err
	}

	return decodedCredentials, nil
}

func SheetSubscribeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	credentials, err := getCredentials()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	client, err := firestore.NewClient(ctx, "drink-and-eat-b7e64", option.WithCredentialsJSON(credentials))
	if err != nil {
		logError("Failed to create Firestore client: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()
	// Check if the Firestore client is successfully initialized by performing a simple operation
	cols, err := client.Collections(ctx).GetAll()
	if err != nil {
		logError("Failed to verify Firestore client initialization: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Firebase app created successfully. Collection path: %d", len(cols))
	if err := json.NewEncoder(w).Encode(map[string]int{
		"collections-count": len(cols),
	}); err != nil {
		logError("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func logError(format string, v ...interface{}) {
	err := fmt.Errorf(format, v...)
	log.Printf("[Error] %s", err)
}
