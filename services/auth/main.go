package main

import (
	"log"
	"net/http"
)

const AUTH_ADDR string = ":5010"

func main() {
	g := NewAuthService(AUTH_ADDR)

	Db = NewDatabase()
	SeedDatabase()

	log.Printf("Auth Service running on port %s", AUTH_ADDR)
	err := http.ListenAndServe(AUTH_ADDR, g.Mux)
	if err != nil {
		log.Fatal(err)
	}
}
