package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	serverType := os.Getenv("SERVER_TYPE")

	switch strings.ToUpper(serverType) {
	case "HTTP":
		log.Println("Starting HTTP server")
	case "GRPC":
		log.Println("Starting gRPC server")
	default:
		log.Fatalf("Unknown server type: %s\n", serverType)
	}
}
