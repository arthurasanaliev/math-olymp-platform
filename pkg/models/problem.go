package models

// Problem represents a problem in the database
type Problem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	Tags       string `json:"tags"`
	Statement  string `json:"statement"`
	Tutorial   string `json:"tutorial"`
	Answer     int    `json:"answer"`
}
