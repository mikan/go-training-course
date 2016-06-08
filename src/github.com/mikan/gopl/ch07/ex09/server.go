// Copyright 2016 mikan. All rights reserved.

package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

var itemList = template.Must(template.New("itemlist").Parse(`
<html>
<head>
<title>CH07-EX09</title>
<style>
body {
    background-color: whitesmoke;
}
table, th, td {
    border-collapse: collapse;
    border: 1px solid gray;
    padding: 5px;
}
th {
    background-color: silver;
}
th a {
    color: gray;
    text-decoration: none;
}
th a:hover {
    color: blue;
    text-decoration: underline;
}
tr {
    text-align: left;
}
#error {
    color: red;
}
</style>
</head>
<body>
<h1>{{.TotalCount}} items</h1>
<table>
<tr style='text-align: left'>
  <th>Title <a href="/?by=title">▼</a></th>
  <th>Artist <a href="/?by=artist">▼</a></th>
  <th>Album <a href="/?by=album">▼</a></th>
  <th>Year <a href="/?by=year">▼</a></th>
  <th>Length <a href="/?by=length">▼</a></th>
</tr>
{{range .Items}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
<p id="error">{{.Error}}</p>
</body>
</html>
`))

type ItemResult struct {
	TotalCount int
	Items      []*Track
	Error      string
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var history []tableWidgetSort // Warning: thread shared

func main() {
	history = append(history, tableWidgetSort{tracks, func(x, y *Track) bool { return x.Title < y.Title }, nil})
	history[0].history = history
	http.HandleFunc("/", list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func list(w http.ResponseWriter, req *http.Request) {
	by := req.URL.Query().Get("by")
	var sortBy tableWidgetSort
	err := false
	switch by {
	case "title":
		sortBy = tableWidgetSort{tracks, func(x, y *Track) bool { return x.Title < y.Title }, nil}
	case "artist":
		sortBy = tableWidgetSort{tracks, func(x, y *Track) bool { return x.Artist < y.Artist }, nil}
	case "album":
		sortBy = tableWidgetSort{tracks, func(x, y *Track) bool { return x.Album < y.Album }, nil}
	case "year":
		sortBy = tableWidgetSort{tracks, func(x, y *Track) bool { return x.Year < y.Year }, nil}
	case "length":
		sortBy = tableWidgetSort{tracks, func(x, y *Track) bool { return x.Length < y.Length }, nil}
	default:
		err = true
	}
	if !err {
		history = append(history, sortBy)
		history[len(history)-1].history = history
		sort.Sort(history[len(history)-1])
	}

	var result ItemResult
	result.TotalCount = len(tracks)
	result.Items = tracks
	if err && by != "" {
		result.Error = "Unknown target: " + by
	}
	if err := itemList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
