package main

const (
	cacheFile      = "xkcd.json"
	baseURLPath    = "https://xkcd.com"
	descriptionObj = "info.0.json"
)

//Comic class
type Comic struct {
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Transcript string `json:"transcript"`
	Link       string `json:"link"`
}

var timeIndex map[string]([]Comic) = map[string]([]Comic){}

var numIndex map[int]bool = map[int]bool{}
