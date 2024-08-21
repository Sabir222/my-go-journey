package server

import (
	"github.com/rs/cors"
	"net/http"
	"sabir222/http-calculator/internal/handler"
	"sabir222/http-calculator/internal/middleware"
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

	timerMiddleware := middleware.Time_stam

	return timerMiddleware(corsMiddleware.Handler(mux))
}
