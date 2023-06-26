package main

import (
	"net/http"

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
	// router.Get("/zincsearch", controllers.Zincsearch)

	http.ListenAndServe(":3000", router)
}

// func main() {
// 	r := chi.NewRouter()

// 	// Ruta para el endpoint /indexer
// 	r.Post("/indexer", IndexerHandler)

// 	// Inicia el servidor
// 	http.ListenAndServe(":8080", r)
// }

// router.Use(middleware.Logger)
// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("welcome"))
// })
// http.ListenAndServe(":3000", router)
// services.ProcessEmail()
