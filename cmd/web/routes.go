package main

import (
	"net/http"

	"github.com/Marvellous-Chimaraoke/bookings/pkg/config"
	"github.com/Marvellous-Chimaraoke/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

// routes
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Add some middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
