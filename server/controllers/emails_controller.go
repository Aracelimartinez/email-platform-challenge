package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Aracelimartinez/email-platform-challenge/server/models"
	"github.com/Aracelimartinez/email-platform-challenge/server/services"
	"github.com/Aracelimartinez/email-platform-challenge/server/services/zincsearch"
)

func IndexEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Failed to obtain the usernames"))
		return
	}
	// Configure the maximum number of concurrent workers
	maxWorkers := 8
	workerSem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup
	wg.Add(len(*users))
	var mu sync.Mutex
	var errorOccurred bool
	var errors []error

	for _, user := range *users {
		workerSem <- struct{}{} // Acquire a worker slot

		go func(user string) {
			defer func() {
				<-workerSem // Release the worker slot
				wg.Done()
			}()

			log.Printf("Extracting emails from %s... \n", user)
			userEmails, err := services.ExtractEmailsByUser(user)
			if err != nil {
				mu.Lock()
				errorOccurred = true
				errors = append(errors, fmt.Errorf("failed to extract emails for user %s: %w", user, err))
				mu.Unlock()
				return
			}

			log.Printf("Indexing emails from %s \n", user)
			res, err := zincsearch.CreateDocument("testeemails", userEmails)
			if err != nil {
				mu.Lock()
				errorOccurred = true
				errors = append(errors, fmt.Errorf("failed to index emails for user %s: %w", user, err))
				mu.Unlock()
				return
			}

			log.Printf("Indexed %d documents from user: %s\n", res.RecordCount, user)
		}(user)
	}
	wg.Wait()

	if errorOccurred {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to process emails"))
		for _, err := range errors {
			log.Println(err)
		}
		return
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
