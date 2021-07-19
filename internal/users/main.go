package main

import (
	"log"
	"os"
	"strings"

	"github.com/TofuOverdose/pi-user-service/internal/users/ports"
)

func main() {
	serverType := os.Getenv("SERVER_TYPE")

	app := makeApp()
	switch strings.ToUpper(serverType) {
	case "HTTP":
		log.Println("Starting HTTP server")
		log.Fatal(ports.ServeHttp(":"+getEnvString("HTTP_PORT"), app))
	case "GRPC":
		log.Println("Starting gRPC server")
	default:
		log.Fatalf("Unknown server type: %s\n", serverType)
	}
}
