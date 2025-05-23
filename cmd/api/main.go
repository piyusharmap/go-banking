package main

import "github.com/piyusharmap/go-banking/internal/server"

func main() {
	server := server.NewAPIServer(":8080")

	server.Run()
}
