package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rhenerp/go-htmx-bootstrap/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	
	router.Get("/", handlers.Home)


	log.Printf("Server Started PORT :3000")
	error := http.ListenAndServe(":3000", router)
	if error != nil {
		log.Fatalf("Failed to start server")
	}
}

