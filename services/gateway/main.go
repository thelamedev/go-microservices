package main

import (
	"log"
	"net/http"
)

const GATEWAY_ADDR string = ":5000"

func main() {
	g := NewGateway(GATEWAY_ADDR)

	// Register Services
	g.ServiceRegistry.AddService(registryItem{
		ServiceName:    "auth",
		ServiceBaseUrl: "http://localhost:5010",
		HealthUrl:      "http://localhost:5010/health",
		Health:         1.0,
		TimeFrame:      100,
	})

	log.Printf("Gateway running on port %s", GATEWAY_ADDR)
	err := http.ListenAndServe(GATEWAY_ADDR, g.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
