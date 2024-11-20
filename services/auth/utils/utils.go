package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	WriteJSON(w, map[string]any{
		"error":      err.Error(),
		"statusCode": statusCode,
	}, statusCode)
}

func WriteJSON(w http.ResponseWriter, body map[string]any, statusCode int) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func NewId() string {
	var id []byte
	var chars []byte = []byte("abcdef0123456789")
	for i := 0; i < 16; i++ {
		x := rand.Intn(16)
		id = append(id, chars[x])
	}

	return string(id)
}
