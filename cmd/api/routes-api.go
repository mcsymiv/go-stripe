package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (a *app) routes() http.Handler {
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

	mux.Get("/api/payment-intent", a.GetPaymentIntent)

	return mux
}
