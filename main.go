package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/rhenerp/go-htmx-bootstrap/assets"
	"github.com/rhenerp/go-htmx-bootstrap/config"
	"github.com/rhenerp/go-htmx-bootstrap/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing anyway...")
	}

	cfg := config.LoadConfig()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	var assetsRoot http.FileSystem
	if (cfg.ENV == "development") {
		assetsRoot = http.Dir("./assets")
	} else {
		assetsRoot = http.FS(assets.Assets)
	}
	fmt.Println("Static file served at:", assetsRoot)
	assetsHandler(router, "/assets", assetsRoot, cfg)

	router.Get("/", handlers.Home)
	router.Get("/items", handlers.GetItems)
	router.Get("/items/{itemId}", handlers.GetItemById)

	log.Println("App running in", cfg.ENV, "on PORT:", cfg.Port)

	error := http.ListenAndServe(":" + cfg.Port, router)
	if error != nil {
		log.Fatalf("Failed to start server")
	}
}

func assetsHandler(r chi.Router, path string, root http.FileSystem, cfg config.Config){
	// TODO: Use embed in release ENV
	if path == "" {
		panic("FileServer path cannot be empty")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	
	r.Get(path+"/*", func(w http.ResponseWriter, r *http.Request) {
		
		if (cfg.ENV == "development"){
			w.Header().Set("Cache-Control", "no-store")
		}
		fs.ServeHTTP(w, r)
	})
}