package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/models"
	"github.com/Aracelimartinez/email-platform-challenge/server/services"
	"github.com/Aracelimartinez/email-platform-challenge/server/services/zincsearch"
)

func IndexEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error

	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Failed to obtain the usernames"))
		return
	}

	for _, user := range *users {

		log.Printf("Extracting emails from %s... \n", user)
		userEmails, err := services.ExtractEmailsByUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write([]byte("Failed to extract the emails"))
			return
		}

		log.Printf("Indexing emails from %s \n", user)
		res, err := zincsearch.CreateDocument(models.EmailIndexName, userEmails)
		if err != nil {
			log.Printf("error indexing emails from user: %s ...\n", userEmails)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write([]byte("Error indexing emails to zincsearch API"))
			return
		}

		log.Printf("Indexed %d documents from user: %s\n", res.RecordCount, user)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Emails indexed succesfully"))
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	query := r.URL.Query().Get("query")
	from := 0
	maxResults := 15

	searchResponse, err := zincsearch.SearchDocuments(models.EmailIndexName, query, from, maxResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error searching into emails at the zincsearch API"))
		return
	}

	emails, err := services.MapZincSearchEmails(searchResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Failed mapping the emails"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emails)
}
