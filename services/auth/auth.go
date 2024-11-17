package main

import (
	"net/http"

	"github.com/thelamedev/go-microservices-auth/utils"
)

type AuthService struct {
	Addr string

	Mux *http.ServeMux
}

func NewAuthService(addr string) *AuthService {
	mux := http.NewServeMux()

	g := &AuthService{
		Addr: addr,
		Mux:  mux,
	}

	mux.HandleFunc("POST /login", NoOpHandle)

	return g
}

func NoOpHandle(w http.ResponseWriter, r *http.Request) {
	req_path := r.URL.Path

	utils.WriteJSON(w, map[string]any{
		"message": "No Operation.",
		"path":    req_path,
	}, 200)
}
