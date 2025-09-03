package main

import (
	"log"

	router "github.com/SLANGERES/Service-Discovery/internal/Router"
)

func StartServer(addr string) {
	r := router.RegisterRouter()
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
func main() {
	StartServer(":9000")
}
