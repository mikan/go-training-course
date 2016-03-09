// Copyright 2016 mikan. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/mikan/util/crypt"
	"github.com/mikan/util/input"
)

const searchURL = "http://www.omdbapi.com/?r=json&t="

var encryptedAPIKey = []byte{85, 94, 58, 196, 151, 219, 205, 244, 117, 22, 2, 102, 91, 79, 123, 250}

func main() {
	// enc,_ := crypt.Encrypt("API KEY", "PASSWORD")
	// fmt.Printf("%d\n", enc)
	var apiKey string
	for {
		password := input.Word("Input Aikotoba")
		if input.IsQuit(password) {
			return
		}
		apiKey, _ = crypt.Decrypt(encryptedAPIKey, password)
		if len(apiKey) == 8 {
			break // correct
		}
		fmt.Println("Illegal Aikotoba: " + password)
	}

	// Main loop
	for {
		query := input.SingleLine("Input query")
		if input.IsQuit(query) {
			return
		}
		list, err := search(query)
		if err != nil {
			log.Fatal(err)
		}
		if list == nil {
			fmt.Println("No such movie: " + query)
			continue
		}
		m := list[0]
		fmt.Println("Found: " + m.Title + " (" + m.Year + ")")

		// Poster
		if input.Word("Donload a poster? {y,n}") != "y" {
			continue
		}
		fmt.Println("Oops! Not implemented yet.")
	}
}

func search(query string) ([]Movie, error) {
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
