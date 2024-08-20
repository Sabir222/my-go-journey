package server

import (
	"net/http"
	"sabir222/http-calculator/internal/handler"
)

func RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloWorldHandler)
	mux.HandleFunc("/user", handler.UserHandler)
	mux.HandleFunc("/req", handler.SubstractionHandler)
	return mux
}
