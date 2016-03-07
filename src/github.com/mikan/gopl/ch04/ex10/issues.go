// Copyright 2016 mikan. All rights reserved.

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mikan/gopl/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	dMonth, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*30))
	dYear, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*365))
	var lessThanAMonth []*github.Issue
	var lessThanAYearOld []*github.Issue
	var moreThanAYearOld []*github.Issue

	// Classify issues by 3 durations
	for _, item := range result.Items {
		if item.CreatedAt.After(now.Add(-dMonth)) {
			lessThanAMonth = append(lessThanAMonth, item)
		} else if item.CreatedAt.After(now.Add(-dYear)) {
			lessThanAYearOld = append(lessThanAYearOld, item)
		} else {
			moreThanAYearOld = append(moreThanAYearOld, item)
		}
	}

	// Print each collections
	fmt.Printf("%d issues in less than a month:\n", result.TotalCount)
	for _, item := range lessThanAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues in less than a year old:\n", result.TotalCount)
	for _, item := range lessThanAYearOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues in more than a year old:\n", result.TotalCount)
	for _, item := range moreThanAYearOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
