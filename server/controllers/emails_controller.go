package controllers

import (
	"log"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/services"
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

	for _, user := range users {

		log.Printf("Extracting emails from %s... \n", user)
		userEmails, err := services.ExtractEmailsByUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			w.Write([]byte("Failed to extract the emails"))
			return
		}

		log.Printf("Indexing emails from %s \n", user)
		res, err := services.CreateDocument("emails", userEmails)
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
