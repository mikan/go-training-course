// Copyright 2016 mikan. All rights reserved.

package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func Search(query string) ([]Movie, error) {
	// Retrieve
	q := url.QueryEscape(query)
	resp, err := http.Get(searchURL + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Println("URL: " + searchURL + q)
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// Parse
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	if movie.Response != "True" {
		return nil, nil
	}
	return []Movie{movie}, nil
}
