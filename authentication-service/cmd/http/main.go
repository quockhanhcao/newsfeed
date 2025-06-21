package main

import (
	"authentication/internal/handler/http"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server, err := http.New()
	if err != nil {
		fmt.Println("error init http server: " + err.Error())
	}
	port := os.Getenv("PORT")
	if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
        fmt.Println("error starting http server: " + err.Error())
    }
}
