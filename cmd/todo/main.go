package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"jk-todolist/internal/server"
	"log"
	"os"
)

func main() {
	// load from .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	r := server.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
