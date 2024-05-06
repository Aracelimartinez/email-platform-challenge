package zincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/config"
)

const (
	zincSearchHost = "http://zincsearch:4080"
)

// CreateDocuments creates documents with the Bulkv2 ZincSearch API
func CreateDocument(index string, records interface{}) (*IndexDocumentsResponse, error) {
	path := "/api/_bulkv2"
	body := IndexDocumentsRequest{
		Index:   index,
		Records: records,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the body into JSON: %w\n", err)
	}

	request, err := http.NewRequest(http.MethodPost, zincSearchHost+path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create the POST request to the ZincSearch API: %w\n", err)
	}

	request.SetBasicAuth(config.ZincSearchCredentials.Username, config.ZincSearchCredentials.Password)
	request.Header.Set("Content-Type", "application/json")

	bodyResponse, err := executeAndReadResponse(request)
	if err != nil {
		return nil, err
	}

	var indexDocumentsResponse IndexDocumentsResponse

	err = json.Unmarshal(bodyResponse, &indexDocumentsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding sucess request response: %w\n", err)
	}

	return &indexDocumentsResponse, nil
}

// executeAndReadResponse read the response from zincsearch
func executeAndReadResponse(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send the POST request to the ZincSearch API: %w\n", err)
	}
	defer response.Body.Close()

	var zincSearchErrorReponse ZincSearchErrorReponse

	bodyResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read the response of the request: %w\n", err)
	}

	if response.StatusCode != http.StatusOK {

		err = json.Unmarshal(bodyResponse, &zincSearchErrorReponse)
		if err != nil {
			return nil, fmt.Errorf("error decoding failed request response: %w\n", err)
		}

		return nil, fmt.Errorf("error from ZincSearch API: %s\n error status code: %d", zincSearchErrorReponse.Error, response.StatusCode)
	}

	return bodyResponse, nil
}

// SearchDocuments search the documents that match with the term given in the zincsearch API
func SearchDocuments(indexName, term string, from, max int) (*SearchDocumentsRsponse, error) {
	path := "/api/" + indexName + "/_search"
	body := setSearchDocumentsRequest(indexName, term, from, max)

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the body into JSON: %w\n", err)
	}

	request, err := http.NewRequest(http.MethodPost, zincSearchHost+path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create the POST request to the ZincSearch API: %w\n", err)
	}
	request.SetBasicAuth(config.ZincSearchCredentials.Username, config.ZincSearchCredentials.Password)
	request.Header.Set("Content-Type", "application/json")

	bodyResponse, err := executeAndReadResponse(request)
	if err != nil {
		return nil, err
	}

	var searchDocumentsResponse SearchDocumentsRsponse

	err = json.Unmarshal(bodyResponse, &searchDocumentsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding searching request response: %w\n", err)
	}

	fmt.Print(searchDocumentsResponse.Hits.Hits)

	return &searchDocumentsResponse, nil
}

// setSearchDocumentsRequest set the respuest body for the SearchDocuments function
func setSearchDocumentsRequest(indexName, term string, from, maxResults int) (body *SearchDocumentsRequest) {
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
