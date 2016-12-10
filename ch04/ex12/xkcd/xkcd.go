// Copyright 2016 mikan. All rights reserved.

package xkcd

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Transcript string // in Markdown format
	SafeTitle  string `json:"safe_title"`
	Alt        string
	Image      string `json:"img"`
	Day        string
}
