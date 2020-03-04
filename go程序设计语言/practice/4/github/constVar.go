package github

import "time"

//IssuesURL url
const IssuesURL = "https://api.github.com/search/issues"

//IssueSearchResult class
type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

//User class
type User struct {
	Login   string
	HTMLURL string `josn:"html_url"`
}

//Issue class
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"create_at"`
	Body      string    //Markdown格式
}
