package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getThirdPartyResponse(url string) map[string]any {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}

	// add the headers
	req.Header.Add("Content-type", "application/json; charset=utf-8")

	// build the client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to do HTTP request %v", err)
	}

	// read the response body from the server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	defer resp.Body.Close()

	data := make(map[string]interface{})
	//days := make(map[string]interface{})

	json.Unmarshal(body, &data)

	return data
}
