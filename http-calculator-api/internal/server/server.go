package server

import (
	"net/http"
	"time"
)

func NewServer() *http.Server {
	mux := RegisterRoutes()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}
