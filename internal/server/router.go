package server

import (
	"fmt"
	"net/http"

	"study-app/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRoutes configures all HTTP routes for the application
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/users", handlers.SignupHandler)
	r.Post("/api/login", handlers.LoginHandler)

	r.Get("/api/decks", handlers.GetDecksHandler)
	r.Post("/api/decks", handlers.CreateDeckHandler)

	r.Get("/api/decks/{id}", handlers.GetDeckHandler)

	r.Post("/api/questions", handlers.AddQuestionHandler)
	r.Get("/api/questions", handlers.ListQuestionsByDeckHandler)

	r.Post("/api/submit", handlers.SubmitAnswerHandler)

	return r
}

// StartServer starts the HTTP server on the specified port
func StartServer(port string) error {
	router := SetupRoutes()

	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}
