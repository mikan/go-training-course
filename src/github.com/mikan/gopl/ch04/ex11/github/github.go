// Copyright 2016 mikan. All rights reserved.

// Package github provies a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

const RepoPrefix = "repo:"
const RepoParam = ":repo"
const NumberParam = ":number"
const IssuesSearchURL = "https://api.github.com/search/issues"
const IssueCreateURL = "https://api.github.com/repos/" + RepoParam + "/issues"
const IssueURL = IssueCreateURL + "/" + NumberParam

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type IssueRequest struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Label     []string `json:"label,omitempty"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Assignee  *User
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
