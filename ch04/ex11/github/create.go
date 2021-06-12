// Copyright 2016 mikan. All rights reserved.

package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// CreateIssue creates a specified issue.
func CreateIssue(repo string, cred *Credential, create *IssueRequest) error {
	url := IssueCreateURL
	url = strings.Replace(url, RepoParam, repo, 1)

	json, _ := json.Marshal(create)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Add(BasicAuth(cred))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("[DEBUG] URL=" + url)
		fmt.Println("[DEBUG] JSON=" + string(json))
		return errors.New("Failed to create: " + resp.Status)
	}
	return nil
}
