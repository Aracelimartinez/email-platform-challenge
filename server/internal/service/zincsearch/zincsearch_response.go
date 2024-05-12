package zincsearch

// ZincSaerchErrorReponse is the response of ZincSearch when an error occurs
type ZincSearchErrorReponse struct {
	Error string `json:"error"`
}

// IndexDocumentsResponse is the response of ZincSearch when a document is created
type IndexDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// SearchDocumentsRsponse is the response of ZincSearch when search for documents
type SearchDocumentsRsponse struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	MaxScore float64     `json:"max_score"`
	Hits     Hits        `json:"hits"`
	Buckets  interface{} `json:"buckets"`
	Error    string      `json:"error"`
}

// Simplified SearchResponse structure
type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type Hit struct {
	Index     string                 `json:"_index"`
	Type      string                 `json:"_type"`
	ID        string                 `json:"_id"`
	Score     float64                `json:"_score"`
	Timestamp string                 `json:"@timestamp"`
	Source    map[string]interface{} `json:"_source"`
}

type Total struct {
	Value int `json:"value"`
}
