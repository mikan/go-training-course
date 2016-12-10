// Copyright 2016 mikan. All rights reserved.

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const milestonesURL = "https://api.github.com/repos/" + RepoParam + "/milestones"

// GetMilestones gets a list of milestones.
func GetMilestones(repo string) ([]Milestone, error) {
	url := strings.Replace(milestonesURL, RepoParam, repo, 1)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
