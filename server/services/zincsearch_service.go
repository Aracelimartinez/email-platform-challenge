package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Aracelimartinez/email-platform-challenge/server/config"
	"github.com/Aracelimartinez/email-platform-challenge/server/models"
)

const (
	zincSearchHost = "http://localhost:4080"
)

func CreateDocument(index string, records interface{}) (*models.IndexDocumentsResponse, error) {
	var err error

	path := "/api/_bulkv2"
	body := models.IndexDocumentsRequest{
		Index:   index,
		Records: records,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to convert the body into JSON: %w\n", err)
	}

	request, err := http.NewRequest(http.MethodPost, zincSearchHost+path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to creare the POST request to the ZincSearch API: %w\n", err)
	}

	request.SetBasicAuth(config.ZincSearchCredentials.Username, config.ZincSearchCredentials.Password)
	request.Header.Set("Content-Type", "application/json")

	response, err := executeAndReadResponse(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func executeAndReadResponse(request *http.Request) (*models.IndexDocumentsResponse, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send the POST request to the ZincSearch API: %w\n", err)
	}
	defer response.Body.Close()

	var indexDocumentsResponse models.IndexDocumentsResponse
	var zincSaerchErrorReponse models.ZincSaerchErrorReponse

	bodyResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read the response of the request: %w\n", err)
	}

	if response.StatusCode != http.StatusOK {

		err = json.Unmarshal(bodyResponse, &zincSaerchErrorReponse)
		if err != nil {
			return nil, fmt.Errorf("error decoding failed request response: %w\n", err)
		}

		return nil, fmt.Errorf("error creating documents: %s\n error status code: %d", zincSaerchErrorReponse.Error, err, response.StatusCode)
	}

	err = json.Unmarshal(bodyResponse, &indexDocumentsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding sucess request response: %w\n", err)
	}

	return &indexDocumentsResponse, nil
}
