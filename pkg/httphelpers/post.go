package httphelpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Post data as JSON to the URL
func Post(url string, data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal %v into JSON format\n", data)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Failed to POST data to `%s`. Data: %v. Error: %v", url, data, err)
		return nil, err

	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body. Error: %v", err)
		return nil, err
	}

	return responseBody, nil
}
