package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Aracelimartinez/email-platform-challenge/server/controllers"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	router.Get("/indexer", controllers.IndexEmails)
	router.Get("/search", controllers.SearchEmails)

	//Profiling routes
	router.Mount("/debug", middleware.Profiler())


	http.ListenAndServe(":3000", router)
}
