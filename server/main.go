package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/Aracelimartinez/email-platform-challenge/server/controllers"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Cors config
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://web:5173"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
	router.Use(corsOptions.Handler)

	//Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!!"))
	})
	router.Get("/indexer", controllers.IndexEmails)
	router.Get("/search", controllers.SearchEmails)

	//Profiling routes
	router.Mount("/debug", middleware.Profiler())

	http.ListenAndServe(":3000", router)
}
