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

// EditIssue edits a specified issue.
func EditIssue(repo, number string, cred *Credential, edit *IssueRequest) error {
	url := IssueURL
	url = strings.Replace(url, RepoParam, repo, 1)
	url = strings.Replace(url, NumberParam, number, 1)

	json, _ := json.Marshal(edit)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Add(basicAuth(cred))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("[DEBUG] URL=" + url)
		fmt.Println("[DEBUG] JSON=" + string(json))
		return errors.New("Failed to update: " + resp.Status)
	}
	return nil
}
