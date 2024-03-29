// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"log"

	"io/ioutil"
	"os"

	"net/http"

	"github.com/mikan/go-training-course/ch04/ex13/omdb"
)

var encryptedAPIKey = []byte{46, 143, 219, 179, 208, 231, 113, 28, 238, 20, 156, 16, 11, 143, 167, 225}
var nilMovie omdb.Movie

func main() {
	// To create encryptedAPIKey, uncomment and update with your API key and password (strings import needed).
	// enc,_ := crypt.Encrypt("APIKEY", "PASSWORD")
	// println("encryptedAPIKey=" + strings.Replace(fmt.Sprintf("%d", enc), " ", ",", -1))

	var apiKey = ""

	// Main loop
	for {
		query := SingleLine("Input query")
		if IsQuit(query) {
			return
		}
		m := handleSearch(query)
		if m == nilMovie {
			fmt.Println("No such movie: " + query)
			continue
		}
		fmt.Println("Found: " + m.Title + " (" + m.Year + ")")
		if Word("Donload a poster? {y,n}") == "y" {
			if Word("Select API {1,2} (1=Poster API, 2=Poster Element)") == "1" {
				if apiKey == "" {
					apiKey = handleAPIKeyInput()
					if apiKey == "" {
						return
					}
				}
				handlePoster(m.IMDBID, apiKey)
			} else {
				handlePoster2(m.IMDBID, m.Poster)
			}
		}
	}
}

func handleSearch(query string) omdb.Movie {
	list, err := omdb.Search(query)
	if err != nil {
		log.Fatal(err)
	}
	if list == nil {
		return nilMovie
	}
	return list[0]
}

func handleAPIKeyInput() string {
	for {
		password := Word("Input Aikotoba")
		if IsQuit(password) {
			return ""
		}
		apiKey, _ := Decrypt(encryptedAPIKey, password)
		if len(apiKey) == 8 {
			return apiKey // correct
		}
		fmt.Println("Illegal Aikotoba: " + password)
	}
}

// Fetch poster from Poster API
func handlePoster(id, apiKey string) {
	path := os.TempDir()
	pathInput := SingleLine("Save to [" + path + "]")
	if pathInput != "" {
		path = pathInput
	}
	path += "/" + id
	poster, err := omdb.GetPoster(apiKey, id)
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile(path, poster, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Successfully wrote to " + path)
}

// Fetch poster from element of search result.
func handlePoster2(id, poster string) {
	if poster == "N/A" {
		fmt.Println("Poster is not available.")
		return
	}
	path := os.TempDir()
	pathInput := SingleLine("Save to [" + path + "]")
	if pathInput != "" {
		path = pathInput
	}
	path += "/" + id
	resp, err := http.Get(poster)
	if err != nil {
		log.Fatal(err)
	}
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	err3 := ioutil.WriteFile(path, data, 0644)
	if err3 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Successfully wrote to " + path)
}
