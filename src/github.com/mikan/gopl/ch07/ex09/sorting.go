// Copyright 2016 mikan. All rights reserved.

package main

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type tableWidgetSort struct {
	t       []*Track
	less    func(x, y *Track) bool
	history []tableWidgetSort
}

func (x tableWidgetSort) Len() int { return len(x.t) }
func (x tableWidgetSort) Less(i, j int) bool {
	l := false
	for _, s := range x.history {
		l = s.less(x.t[i], x.t[j])
	}
	return l
}
func (x tableWidgetSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
