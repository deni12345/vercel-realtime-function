package main

import (
	"fmt"
	"github/vercel-realtime-function/configs"
	"github/vercel-realtime-function/utils"
	"log"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Project ID: %v", utils.ToString(cfg))
}
