package main

import (
	"log"
	"net/http"
	"os"

	"epam.com/web-services/library-management/login-service/handlers"
)

func main() {
	handlers.SetupRoutes()
	err := http.ListenAndServe(":"+os.Getenv("Port"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
