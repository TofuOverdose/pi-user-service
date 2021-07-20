package main

import (
	"log"
	"os"
	"strings"

	"github.com/TofuOverdose/pi-user-service/internal/users/ports/grpc"
	"github.com/TofuOverdose/pi-user-service/internal/users/ports/http"
)

func main() {
	serverType := os.Getenv("SERVER_TYPE")

	app := makeApp()
	switch strings.ToUpper(serverType) {
	case "HTTP":
		config := http.ServerConfig{
			Port: ":" + getEnvString("HTTP_PORT"),
		}
		log.Println("Starting HTTP server")
		log.Fatal(http.Serve(app, config))
	case "GRPC":
		config := grpc.ServerConfig{
			Port:          ":" + getEnvString("GRPC_PORT"),
			UseReflection: getEnvString("ENV") == "dev",
		}
		log.Println("Starting gRPC server")
		log.Fatal(grpc.Serve(app, config))
	default:
		log.Fatalf("Unknown server type: %s\n", serverType)
	}
}
