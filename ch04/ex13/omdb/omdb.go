// Copyright 2016 mikan. All rights reserved.

package omdb

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	MetaScore  string `json:"Metascore"`
	IMDBRating string `json:"imdbRating"`
	IMDBVotes  string `json:"imdbVotes"`
	IMDBID     string `json:"imdbID"`
	Type       string
	Response   string
}

const apiKeyParam = "********"
const searchURL = "http://www.omdbapi.com/?r=json&t="
const posterURL = "http://img.omdbapi.com/?apikey=" + apiKeyParam + "&i="
