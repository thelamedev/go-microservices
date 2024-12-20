package main

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/thelamedev/go-microservices-gateway/utils"
)

type Gateway struct {
	Addr string

	Mux             *http.ServeMux
	ServiceRegistry *ServiceRegistry
}

func NewGateway(addr string) *Gateway {
	mux := http.NewServeMux()

	g := &Gateway{
		Addr:            addr,
		Mux:             mux,
		ServiceRegistry: NewServiceRegistry(),
	}

	mux.HandleFunc("/auth/{path}", g.ProxyServiceRequest("auth"))

	return g
}

func (g *Gateway) ProxyServiceRequest(serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")

		reqMethod := r.Method
		serviceItem, err := g.ServiceRegistry.GetService(serviceName)
		if err != nil {
			utils.WriteError(w, err, 500)
			return
		}

		reqUrl, err := url.JoinPath(serviceItem.ServiceBaseUrl, path)
		if err != nil {
			utils.WriteError(w, err, 500)
			return
		}

		newReq, err := http.NewRequest(reqMethod, reqUrl, r.Body)
		defer r.Body.Close()

		newReq.Header.Set("authorization", r.Header.Get("authorization"))
		newReq.Header.Set("origin", r.Header.Get("origin"))

		if err != nil {
			utils.WriteError(w, err, 500)
			return
		}

		response, err := http.DefaultClient.Do(newReq)
		if err != nil {
			utils.WriteError(w, err, 500)
			return
		}

		w.WriteHeader(response.StatusCode)

		defer response.Body.Close()
		_, err = io.Copy(w, response.Body)
		if err != nil {
			log.Printf("Error writing response body to stream: %v", err)
		}
	}
}

func NoOpHandle(w http.ResponseWriter, r *http.Request) {
	req_path := r.URL.Path

	utils.WriteJSON(w, map[string]any{
		"message": "No Operation.",
		"path":    req_path,
	}, 200)
}
