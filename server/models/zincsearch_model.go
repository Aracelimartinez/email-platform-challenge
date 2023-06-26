package models

// ZincSaerchErrorReponse is the response of ZincSearch when an error occurs
type ZincSaerchErrorReponse struct {
	Error  string `json:"error"`
}

// IndexDocumentsRequest is the body request with the document info to be created in Zincsearch
type IndexDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

// IndexDocumentsResponse is the response of ZincSearch when a document is created
type IndexDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}
