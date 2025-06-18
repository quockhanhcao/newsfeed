package main

import (
	"epk14/newsfeed/internal/handler/http"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server, err := http.New()
	if err != nil {
		fmt.Println("error init http server", err)
	}
	port := os.Getenv("PORT")
	server.Start(fmt.Sprintf(":%s", port))
}
