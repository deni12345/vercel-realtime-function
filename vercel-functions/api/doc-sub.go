package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func getCredentials() ([]byte, error) {
	googleCredentials := os.Getenv("GOOGLE_CREDENTIALS")
	if googleCredentials == "" {
		return nil, logError("GOOGLE_CREDENTIALS environment variable is not set")
	}

	decodedCredentials, err := base64.RawStdEncoding.DecodeString(googleCredentials)
	if err != nil {
		return nil, logError("Failed to decode Google credentials: %v", err)
	}

	return decodedCredentials, nil
}

func CollectionSubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	credentials, err := getCredentials()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(credentials))
	if err != nil {
		logError("Failed to create Firebase app: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		logError("Failed to create Firestore client: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	log.Printf("Firebase app created successfully. Collection path: %s", client.Collection("users").Path)
	if err := json.NewEncoder(w).Encode("credentials"); err != nil {
		logError("Failed to encode response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func logError(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v...)
	log.Printf("[Error] %s", err)
	return err
}
