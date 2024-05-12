package zincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/configs"
)

// CreateDocuments creates documents with the Bulkv2 ZincSearch API
func CreateDocument(index string, records interface{}) (*IndexDocumentsResponse, error) {
	path := configs.GlobalConfig.ZincSearchCredentials.Host + "/api/_bulkv2"
	body := IndexDocumentsRequest{
		Index:   index,
		Records: records,
	}

	request, err := newJSONRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	bodyResponse, err := executeAndReadResponse(request)
	if err != nil {
		return nil, err
	}

	var indexDocumentsResponse IndexDocumentsResponse

	err = json.Unmarshal(bodyResponse, &indexDocumentsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding sucess request response: %w", err)
	}

	return &indexDocumentsResponse, nil
}

// SearchDocuments search the documents that match with the term given in the zincsearch API
func SearchDocuments(indexName, term string, from, max int) (*SearchDocumentsRsponse, error) {
	path := configs.GlobalConfig.ZincSearchCredentials.Host + "/api/" + indexName + "/_search"
	body := setSearchDocumentsRequest(term, from, max)

	request, err := newJSONRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}

	bodyResponse, err := executeAndReadResponse(request)
	if err != nil {
		return nil, err
	}

	var searchDocumentsResponse SearchDocumentsRsponse

	err = json.Unmarshal(bodyResponse, &searchDocumentsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding searching request response: %w", err)
	}

	return &searchDocumentsResponse, nil
}

// newJSONRequest prepares a new HTTP request with JSON headers
func newJSONRequest(method, url string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the body into JSON: %w", err)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.SetBasicAuth(configs.GlobalConfig.ZincSearchCredentials.Username, configs.GlobalConfig.ZincSearchCredentials.Password)
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

// executeAndReadResponse read the response from zincsearch
func executeAndReadResponse(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send the POST request to the ZincSearch API: %w", err)
	}
	defer response.Body.Close()

	var zincSearchErrorReponse ZincSearchErrorReponse

	bodyResponse, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read the response of the request: %w", err)
	}

	if response.StatusCode != http.StatusOK {

		err = json.Unmarshal(bodyResponse, &zincSearchErrorReponse)
		if err != nil {
			return nil, fmt.Errorf("error decoding failed request response: %w", err)
		}

		return nil, fmt.Errorf("error from ZincSearch API: %s\n error status code: %d", zincSearchErrorReponse.Error, response.StatusCode)
	}

	return bodyResponse, nil
}

// setSearchDocumentsRequest set the respuest body for the SearchDocuments function
func setSearchDocumentsRequest(term string, from, maxResults int) (body *SearchDocumentsRequest) {
	body = &SearchDocumentsRequest{
		Source:     []string{},
		From:       from,
		MaxResults: maxResults,
		Query: SearchDocumentsQuery{
			Term:  term,
			Field: "_all",
		},
		SearchType: "match",
		SortFields: []string{"-@timestamp"},
	}
	return
}
