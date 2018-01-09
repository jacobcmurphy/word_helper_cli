package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func getData(q query) ([]apiResponse, error) {
	resp, reqErr := http.Get(q.String())
	if reqErr != nil {
		fmt.Println("Error fetching data from API", reqErr)
		return nil, reqErr
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Error reading response body", readErr)
		return nil, readErr
	}

	return parseData(body)
}

func parseData(b []byte) ([]apiResponse, error) {
	var responses []apiResponse
	parseErr := json.Unmarshal(b, &responses)
	if parseErr != nil {
		fmt.Println("Could not parse JSON")
		return nil, parseErr
	}

	return responses, nil
}
