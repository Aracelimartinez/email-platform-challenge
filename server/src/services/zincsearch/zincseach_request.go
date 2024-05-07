package zincsearch

// IndexDocumentsRequest is the body request with the document info to be created in Zincsearch
type IndexDocumentsRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

// SearchDocumentsRequest is the body request with the search info to Zincsearch
type SearchDocumentsRequest struct {
	Source     []string             `json:"_source"`
	From       int                  `json:"from"`
	MaxResults int                  `json:"max_results"`
	Query      SearchDocumentsQuery `json:"query"`
	SearchType string               `json:"search_type"`
	SortFields []string             `json:"sort_fields"`
}

// SearchDocumentsQuery is the query for the SearchDocumentsRequest info
type SearchDocumentsQuery struct {
	Field string `json:"field"`
	Term  string `json:"term"`
}
