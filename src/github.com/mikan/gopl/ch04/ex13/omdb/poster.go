// Copyright 2016 mikan. All rights reserved.

package omdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPoster(apiKey, id string) ([]byte, error) {
	// Retrieve
	url := strings.Replace(posterURL, apiKeyParam, apiKey, 1) + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// Read
	poster, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return poster, nil
}
