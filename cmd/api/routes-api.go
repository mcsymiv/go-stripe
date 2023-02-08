package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mcsymiv/go-stripe/internal/api/config"
	"github.com/mcsymiv/go-stripe/internal/api/handlers"
)

func routes(a *config.Application) http.Handler {
	mux := chi.NewMux()

	// CORS chi middleware
	mux.Use(cors.Handler(cors.Options{
		// Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		// Maximum value not ignored by any of major browsers
		MaxAge: 300,
	}))

	mux.Get("/api/payment-intent", handlers.Repo.GetPaymentIntent)

	return mux
}
