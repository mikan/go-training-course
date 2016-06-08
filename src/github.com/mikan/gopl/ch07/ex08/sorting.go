// Copyright 2016 mikan. All rights reserved.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)

	// EX08
	fmt.Println("\nTableWidget(1/3) Title:")
	var history []tableWidgetSort
	history = append(history, tableWidgetSort{tracks, func(x, y *Track) bool { return x.Title < y.Title }, nil})
	history[0].history = history
	sort.Sort(history[0])
	printTracks(tracks)
	fmt.Println("\nTableWidget(2/3) Year:")
	history = append(history, tableWidgetSort{tracks, func(x, y *Track) bool { return x.Year < y.Year }, nil})
	history[len(history)-1].history = history
	sort.Sort(history[len(history)-1])
	printTracks(tracks)
	fmt.Println("\nTableWidget(3/3) Title:")
	history = append(history, tableWidgetSort{tracks, func(x, y *Track) bool { return x.Title < y.Title }, nil})
	history[len(history)-1].history = history
	sort.Sort(history[len(history)-1])
	printTracks(tracks)

	// sort.Stable
	fmt.Println("\nbyYear(sort.Stable):")
	sort.Stable(byYear(tracks))
	printTracks(tracks)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

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

func init() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}
