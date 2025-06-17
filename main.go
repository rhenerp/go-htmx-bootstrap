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

	assetsHandler(router, "/assets", http.Dir("./assets"))
	router.Get("/", handlers.Home)
	router.Get("/items", handlers.GetItems)


	log.Printf("Server Started PORT :3000")
	error := http.ListenAndServe(":3000", router)
	if error != nil {
		log.Fatalf("Failed to start server")
	}
}

func assetsHandler(r chi.Router, path string, root http.FileSystem){
	// TODO: Use embed in release ENV
	if path == "" {
		panic("FileServer path cannot be empty")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	
	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		fs.ServeHTTP(w, r)
	})
}

