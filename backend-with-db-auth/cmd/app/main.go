package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"sabir222/go-backend/internal/server"
	"strconv"
)

func main() {

	server := server.New()
	server.RegisterRoutes()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
