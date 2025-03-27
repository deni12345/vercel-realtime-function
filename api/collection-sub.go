package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	credentials map[string]string = make(map[string]string)
	one         sync.Once
)

// func getCredentials() {
// 	one.Do(func() {
// 		googleCredentials := os.Getenv("GOOGLE_CREDENTIALS")
// 		if googleCredentials == "" {
// 			log.Fatalf("GOOGLE_CREDENTIALS environment variable is not set %s", googleCredentials)
// 		}
// 		decodedCredentials, err := base64.RawStdEncoding.DecodeString(googleCredentials)
// 		if err != nil {
// 			log.Fatal("Failed to decode Google credentials: ", err)
// 		}
// 		if err := json.Unmarshal(decodedCredentials, &credentials); err != nil {
// 			log.Fatal("Failed to unmarshal Google credentials: ", err)
// 		}
// 	})

// }

func CollectionSubHandler(w http.ResponseWriter, r *http.Request) {

	// Handle the request here using the credentials map
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	getCredentials()
	json.NewEncoder(w).Encode(credentials)
}
