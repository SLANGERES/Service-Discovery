package main

import (
	"log"

	router "github.com/SLANGERES/Service-Discovery/internal/Router"
	"github.com/SLANGERES/Service-Discovery/service"
)

func StartServer(addr string) {
	r := router.RegisterRouter()
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
func main() {
	go service.Cleanup()
	go service.HearBeat()
	StartServer(":9000")
}
