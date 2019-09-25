package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"},
		},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}
	fmt.Println(movies)

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marchaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data1, err1 := json.MarshalIndent(movies, "", "    ")
	if err1 != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data1)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data1, &titles); err != nil {
		log.Fatalf("JSON unmarchaling failed: %s", err)
	}
	fmt.Println(titles)

	var actors []struct {Actors []string}
	if err := json.Unmarshal(data1, &actors); err != nil {
		log.Fatalf("JSON unmarchaling failed: %s", err)
	}
	fmt.Println(actors)

}