package main

import (
	"log"

	"github.com/piyusharmap/go-banking/internal/server"

	"github.com/piyusharmap/go-banking/internal/storage"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loding env file")
	}

	store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":8080", store)

	server.Run()
}
