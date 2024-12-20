package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("[ERROR] %s", err.Error())

	WriteJSON(w, map[string]any{
		"error":      err.Error(),
		"statusCode": statusCode,
	}, statusCode)
}

func WriteJSON(w http.ResponseWriter, body map[string]any, statusCode int) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}
