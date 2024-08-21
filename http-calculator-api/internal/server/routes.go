package server

import (
	"github.com/rs/cors"
	"net/http"
	"sabir222/http-calculator/internal/handler"
)

func RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloWorldHandler)
	mux.HandleFunc("/user", handler.UserHandler)
	mux.HandleFunc(" POST /sub", handler.SubstractionHandler)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return corsMiddleware.Handler(mux)
}
