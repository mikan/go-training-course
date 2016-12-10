// Copyright 2016 mikan. All rights reserved.

// Package github provies a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"time"
)

const RepoPrefix = "repo:"
const RepoParam = ":repo"
const OrgParam = ":org"
const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	State        string
	Title        string
	Description  string
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
	// DueOn        time.Time `json:"due_on"`
	Creator *User
}
