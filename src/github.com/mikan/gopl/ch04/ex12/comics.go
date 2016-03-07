// Copyright 2016 mikan. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"strings"

	"github.com/mikan/gopl/ch04/ex12/xkcd"
	"github.com/mikan/util/input"
)

var nilComic xkcd.Comic

func init() {
	nilComic.Num = -1
}

func main() {
	// Check the latest
	last, err := xkcd.FetchLatestNum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("latest: %d\n", last)

	// Fetch all json files
	var files []string
	if input.Word("Do you want to retrive index? {y,n}") == "y" {
		var err error
		files, err = xkcd.FetchAll(last)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		files, _ = xkcd.GetCachedAll()
	}
	fmt.Printf("Found %d items.\n", len(files))

	// Main loop
	for {
		q := input.SingleLine("Query (or \"quit\")")
		if input.IsQuit(q) {
			return
		}
		ch := make(chan xkcd.Comic)
		for _, file := range files {
			go match(file, q, ch)
		}
		for range files {
			printComic(<-ch)
		}
	}
}

func match(file, query string, ch chan<- xkcd.Comic) {
	comic := parse(file)
	if strings.Contains(comic.SafeTitle, query) {
		ch <- comic
	} else if strings.Contains(comic.Transcript, query) {
		ch <- comic
	} else {
		ch <- nilComic
	}
}

func parse(file string) xkcd.Comic {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var comic xkcd.Comic
	if err := json.Unmarshal(data, &comic); err != nil {
		log.Fatal(err)
	}
	return comic
}

func printComic(comic xkcd.Comic) {
	if comic == nilComic {
		return
	}
	fmt.Printf("========== #%d %v\n", comic.Num, comic.SafeTitle)
	fmt.Printf("Title: %v\n", comic.SafeTitle)
	fmt.Printf("URL:   http://xkcd.com/%d/\n", comic.Num)
	fmt.Printf("Transcript: \n%v\n\n", comic.Transcript)
}
