// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"log"

	"io/ioutil"
	"os"

	"github.com/mikan/gopl/ch04/ex13/omdb"
	"github.com/mikan/util/crypt"
	"github.com/mikan/util/input"
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
		query := input.SingleLine("Input query")
		if input.IsQuit(query) {
			return
		}
		m := handleSearch(query)
		if m == nilMovie {
			fmt.Println("No such movie: " + query)
			continue
		}
		fmt.Println("Found: " + m.Title + " (" + m.Year + ")")
		if input.Word("Donload a poster? {y,n}") == "y" {
			if apiKey == "" {
				apiKey = handleAPIKeyInput()
				if apiKey == "" {
					return
				}
			}
			handlePoster(m.IMDBID, apiKey)
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
		password := input.Word("Input Aikotoba")
		if input.IsQuit(password) {
			return ""
		}
		apiKey, _ := crypt.Decrypt(encryptedAPIKey, password)
		if len(apiKey) == 8 {
			return apiKey // correct
		}
		fmt.Println("Illegal Aikotoba: " + password)
	}
}

func handlePoster(id, apiKey string) {
	path := os.TempDir()
	pathInput := input.SingleLine("Save to [" + path + "]")
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
