package router

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/Aracelimartinez/email-platform-challenge/server/configs"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/handler"
)

// Generate returns a router with configured routes.
func Generate() http.Handler {
	router := chi.NewRouter()
	return configRoutes(router)
}

// Configures middlewares, the API and profiler routes.
func configRoutes(router *chi.Mux) http.Handler {
	//CORS Config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{configs.GlobalConfig.FrontEndAdd, "http://localhost:8080"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	router.Use(corsOptions.Handler)
	router.Use(middleware.Logger)

	//Profiling routes
	router.Mount("/debug", middleware.Profiler())

	//API routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	router.Post("/indexer", handler.IndexEmails)
	router.Get("/search", handler.SearchEmails)

	return router
}
