// Copyright 2016 mikan. All rights reserved.

package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
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

func main() {
	http.HandleFunc("/", list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func list(w http.ResponseWriter, req *http.Request) {
	bys := req.URL.Query().Get("by")
	err := false
	var list []lessFunc
	for _, by := range strings.Split(bys, ",") {
		switch by {
		case "title":
			list = append(list, func(c1, c2 *Track) bool { return c1.Title < c2.Title })
		case "artist":
			list = append(list, func(c1, c2 *Track) bool { return c1.Artist < c2.Artist })
		case "album":
			list = append(list, func(c1, c2 *Track) bool { return c1.Album < c2.Album })
		case "year":
			list = append(list, func(c1, c2 *Track) bool { return c1.Year < c2.Year })
		case "length":
			list = append(list, func(c1, c2 *Track) bool { return c1.Length < c2.Length })
		default:
			err = true
		}
	}

	if !err {
		OrderedBy(list).Sort(tracks)
	}

	var result ItemResult
	result.TotalCount = len(tracks)
	result.Items = tracks
	if err && bys != "" {
		result.Error = "Unknown target: " + bys
	}
	if err := itemList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
