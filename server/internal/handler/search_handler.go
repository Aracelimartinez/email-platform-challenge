package handler

import (
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/internal/model"
	"github.com/Aracelimartinez/email-platform-challenge/server/internal/response"
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
		response.Err(w, http.StatusBadRequest, err)
		return
	}

	emails, err := service.MapZincSearchEmails(searchResponse)
	if err != nil {
		response.Err(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, emails)
}
