package main

import (
	"encoding/json"
	"fmt"
	"os"

)

//Movie the movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //如果成员的值为空或者零值，则不输出该值
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca",
		Year:  1942,
		Color: true,
		Actors: []string{
			"huanglin",
			"actor",
		},
	},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Fprintln(os.Stderr, data)
	}
	fmt.Printf("%s\n", data)

	//缩进格式
	data1, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		fmt.Fprintln(os.Stderr, data)
	}
	fmt.Printf("%s\n", data1)

	movie1 := []Movie{}
	err = json.Unmarshal(data1, &movie1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", movie1)
}
