package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wildanfaz/backendgolang2_week8/src/routers"
)

func main() {
	mainRoute, err := routers.New()

	if err != nil {
		log.Fatal(err)
	}

	newErr := godotenv.Load(".env")

	if newErr != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	fmt.Println("Running On Port", port)

	http.ListenAndServe(port, mainRoute)
}
