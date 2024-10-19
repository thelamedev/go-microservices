package main

import (
	"log"
	"net/http"
)

const GATEWAY_ADDR string = ":5000"

func main() {
	g := NewGateway(GATEWAY_ADDR)

	log.Printf("Gateway running on port %s", GATEWAY_ADDR)
	err := http.ListenAndServe(GATEWAY_ADDR, g.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
