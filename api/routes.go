package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *server) routes() {
	s.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	s.mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/status", s.handleStatus())
		r.Post("/url", s.handleShortenURL())
		r.Get("/url/{id}", s.handleGetURL())
		r.Get("/count/{id}", s.handleGetCount())
	})
}
