package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Tatannic",
		year:   1942,
		Color:  true,
		Actors: []string{"Zhang"},
	},
	{
		Title:  "Casablanc",
		year:   1942,
		Color:  false,
		Actors: []string{"Hu", "Zandy"},
	},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatal("JSON marshal failed")
	}
	fmt.Printf("%s", data)
}
