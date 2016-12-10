// Copyright 2016 mikan. All rights reserved.

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetIssue gets a specified issue.
func GetIssue(repo, number string) (*Issue, error) {
	url := IssueURL
	url = strings.Replace(url, RepoParam, repo, 1)
	url = strings.Replace(url, NumberParam, number, 1)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
