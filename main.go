// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahithChigurupati/go-server/models"
	"github.com/MahithChigurupati/go-server/routes"
)

func main() {
	models.InitializeDB()

	router := routes.Router()
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
