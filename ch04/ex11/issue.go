// Copyright 2016 mikan. All rights reserved.

// issue.go provides following issue operations:
// * create
// * read
// * update
// * delete
// * print (out of requirements)
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mikan/go-training-course/ch04/ex11/github"
)

func main() {
	// Setup the target repository
	repo := "mikan/go-training-course"
	if len(os.Args[1:]) > 0 {
		repo = os.Args[1]
		if strings.Contains(repo, "/") {
			log.Fatal("Please specify \"user/repository\" format.")
			return
		}
	}

	// Input login credential
	var cred github.Credential
	cred.Username = Word("Username")
	cred.Password = Word("Password")

	// Main loop
	for {
		switch Word("Input {c,r,u,d,p}") {
		case "c":
			create(repo, &cred)
		case "r":
			read(repo, Word("Input issue num"))
		case "u":
			update(repo, Word("Input issue num"), &cred)
		case "d":
			del(repo, Word("Input issue num"), &cred)
		case "p":
			printIssues(searchIssues([]string{github.RepoPrefix + repo}))
		default:
			return
		}
	}
}

// Handles "c" operation.
func create(repo string, cred *github.Credential) {
	var edit github.IssueRequest
	edit.Title = SingleLine("Input title")
	edit.Body = MultiLine("Input body")

	// Confirmation
	if Word("Are you sure to create? {y,n}") != "y" {
		fmt.Println("Create aborted.")
		return
	}

	// send the request
	err := github.CreateIssue(repo, cred, &edit)
	if err != nil {
		log.Fatal(err)
	}
}

// Handles "r" operation.
func read(repo, id string) {
	issue := getIssue(repo, id)
	printIssue(issue)
}

// Handles "u" operation.
func update(repo, id string, cred *github.Credential) {
	// Retrieve and display the current information
	issue := getIssue(repo, id)
	var assignee string
	if issue.Assignee == nil {
		assignee = "(nobody assigned)"
	} else {
		assignee = issue.Assignee.Login
	}
	fmt.Println("Title:    " + issue.Title)
	fmt.Println("State:    " + issue.State)
	fmt.Println("Body:     " + strings.Replace(issue.Body, "\n", "", -1))
	fmt.Println("Assignee: " + assignee)

	// Editing type selection
	var edit github.IssueRequest
	switch Word("Which do you want to change? {t,s,b,a}") {
	case "t":
		edit.Title = SingleLine("Input title")
	case "s":
		edit.State = Word("Input state {open,closed}")
	case "b":
		edit.Body = MultiLine("Input body")
	case "a":
		edit.Assignee = Word("Input assignee")
	default:
		fmt.Println("Update aboted.")
		return
	}

	// Collect changes
	modTitle := "(unchanged)"
	modState := modTitle
	modBody := modTitle
	modAssignee := modTitle
	if edit.Title != "" {
		modTitle = edit.Title
	}
	if edit.State != "" {
		modState = edit.State
	}
	if edit.Body != "" {
		modBody = edit.Body
	}
	if edit.Assignee != "" {
		modAssignee = edit.Assignee
	}
	fmt.Println("Title:    " + modTitle)
	fmt.Println("State:    " + modState)
	fmt.Println("Body:     " + modBody)
	fmt.Println("Assignee: " + modAssignee)

	// Confirmation
	if Word("Are you sure to change? {y,n}") != "y" {
		fmt.Println("Update aboted.")
		return
	}

	// Send the request
	err := github.EditIssue(repo, id, cred, &edit)
	if err != nil {
		log.Fatal(err)
	}
}

// Handles "d" operation.
func del(repo, id string, cred *github.Credential) {
	// Retrieve & check condition.
	issue := getIssue(repo, id)
	if issue.State == "closed" {
		fmt.Println("Issue #" + strconv.Itoa(issue.Number) + " was already closed.")
		return
	}

	// Confirmation
	if Word("Are you shure to close #"+id+"? {y,n}") != "y" {
		fmt.Println("Close aboted.")
		return
	}

	var edit github.IssueRequest
	edit.State = "closed"

	// Send the request
	err := github.EditIssue(repo, id, cred, &edit)
	if err != nil {
		log.Fatal(err)
	}
}

// Handles "p" operation.
func printIssues(result *github.IssuesSearchResult) {
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func searchIssues(param []string) *github.IssuesSearchResult {
	fmt.Print("Retrieving...")
	result, err := github.SearchIssues(param)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" done.")
	return result
}

func getIssue(repo, id string) *github.Issue {
	fmt.Print("Retrieving...")
	issue, err := github.GetIssue(repo, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" done.")
	return issue
}

func printIssue(item *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s (created at %v) asignee: %s\n",
		item.Number, item.User.Login, item.Title, item.CreatedAt, item.Assignee.Login)
	fmt.Printf("%v\n", item.Body)
}
