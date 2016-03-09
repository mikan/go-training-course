// Copyright 2016 mikan. All rights reserved.

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const usersURL = "https://api.github.com/orgs/" + OrgParam + "/members"

// GetUsers gets a list of members in specified organization.
func GetMembers(org string) ([]User, error) {
	url := strings.Replace(usersURL, OrgParam, org, 1)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
