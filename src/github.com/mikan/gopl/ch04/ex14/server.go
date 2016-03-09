// Copyright 2016 mikan. All rights reserved.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/mikan/gopl/ch04/ex14/github"
)

var issueList = template.Must(template.New("report").Parse(`
<html>
<head>
<title>GitHub Report</title>
<style>
body {
    background-color: whitesmoke;
}
th, td {
    border: 1px solid gray;
}
th {
    background-color: silver;
}
tr {
    style: text-align: left;
}
</style>
</head>
<body>
<h1>GitHub Report</h1>
<h2>{{.Issues.TotalCount}} issues</h2>
<table>
<tr>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues.Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h2>Milestones</h2>
<table>
<tr>
  <th>#</th>
  <th>State</th>
  <th>Open Issues</th>
  <th>Closed Issues</th>
  <th>Title</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{.OpenIssues}}</td>
  <td>{{.ClosedIssues}}</td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
<h2>Users</h2>
<ul>
{{range .Users}}
<li><a href='{{.HTMLURL}}'>{{.Login}}</a></li>
{{end}}
</ul>
</body>
</html>
`))

type Report struct {
	Issues     *github.IssuesSearchResult
	Milestones []github.Milestone
	Users      []github.User
}

func main() {
	// Setup the target repository
	repo := "golang/go"
	if len(os.Args[1:]) > 0 {
		repo = os.Args[1]
		if strings.Contains(repo, "/") {
			log.Fatal("Please specify \"user/repository\" format.")
			return
		}
	}

	// Search bugs
	resp, err := github.SearchIssues("repo:" + repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues loaded.\n", resp.TotalCount)

	// Get milestones
	ms, err := github.GetMilestones(repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d milestones loaded.\n", len(ms))

	// Get collaborators
	users, err := github.GetMembers(repo[:strings.Index(repo, "/")])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d members loaded.\n", len(users))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var report Report
		report.Issues = resp
		report.Milestones = ms
		report.Users = users
		issueList.Execute(w, report)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
