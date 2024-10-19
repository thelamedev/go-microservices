package main

import (
	"net/http"

	"github.com/thelamedev/go-microservices-gateway/utils"
)

type Gateway struct {
	Addr string

	Mux *http.ServeMux
}

func NewGateway(addr string) *Gateway {
	mux := http.NewServeMux()

	g := &Gateway{
		Addr: addr,
		Mux:  mux,
	}

	mux.HandleFunc("POST /auth/login", NoOpHandle)

	return g
}

func NoOpHandle(w http.ResponseWriter, r *http.Request) {
	req_path := r.URL.Path

	utils.WriteJSON(w, map[string]any{
		"message": "No Operation.",
		"path":    req_path,
	}, 200)
}
