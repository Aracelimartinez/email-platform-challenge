package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/internal/model"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/service"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/service/zincsearch"
)

func SearchEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	query := r.URL.Query().Get("query")
	from := 0
	maxResults := 15

	searchResponse, err := zincsearch.SearchDocuments(model.EmailIndexName, query, from, maxResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Write([]byte("Error searching into emails at the zincsearch API"))
		return
	}

	emails, err := service.MapZincSearchEmails(searchResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("Failed mapping the emails"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emails)
}
