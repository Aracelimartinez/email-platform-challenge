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

// curl -X 'POST' \
//   'http://localhost:4080/api/_bulkv2' \
//   -H 'accept: application/json' \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "index": "string",
//   "records": [
//     {
//       "additionalProp1": {}
//     }
//   ]
// }'

// http://localhost:4080/api/_bulkv2

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
	request.Header.Set("Accept", "application/json")

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

// func CreateDocument(userEmail models.UserEmails) ([]byte, error) {
// 	// Create a new product to add to Zincsearch

// 	// Convert the product to JSON
// 	productJSON, err := json.Marshal(userEmail)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Send a POST request to the Zincsearch API to add the new product
// 	request, err := http.NewRequest("POST","http://localhost:4080/api/_bulkv2", bytes.NewBuffer(productJSON))
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.SetBasicAuth("admin", "Complexpass#123")

// 	bytes, err := performRequestAndReadResponse(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes, nil
// }

// func performRequestAndReadResponse(request *http.Request) ([]byte, error) {
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
// 	if !statusOK {
// 		errMsg := fmt.Sprintf("Status code is not OK: %d", response.StatusCode)
// 		return nil, errors.New(errMsg)
// 	}
// 	bytes, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes, nil
// }
