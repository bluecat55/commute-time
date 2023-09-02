package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

type config struct {
	Home   string `json:"home"`
	Office string `json:"office"`
	APIKey string `json:"apiKey"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	home := os.Getenv("HOME_LOC_NAME")
	office := os.Getenv("OFFICE_LOC_NAME")
	apiKey := os.Getenv("API_KEY")

	fmt.Printf("Home: %s\nOffice: %s\n", home, office)

	args := os.Args
	var q *maps.DirectionsRequest
	mode := ""

	if len(args) < 2 {
		fmt.Println("Mode not specified, defaulting to return")
		mode = "return"
	} else if args[1] == "r" || args[1] == "return" {
		mode = "return"
	} else {
		mode = "depart"
	}

	if mode == "return" {
		q = &maps.DirectionsRequest{
			Origin:      office,
			Destination: home,
		}
	} else {
		q = &maps.DirectionsRequest{
			Origin:      home,
			Destination: office,
		}
	}

	fmt.Printf("Mode: %s\n", mode)

	mc, merr := maps.NewClient(maps.WithAPIKey(apiKey))
	if merr != nil {
		log.Fatalln("Error initing map client")
	}

	routes, _, err := mc.Directions(context.Background(), q)
	var tds []time.Duration
	for i := range routes {
		var td time.Duration
		for j := range routes[i].Legs {
			for k := range routes[i].Legs[j].Steps {
				td = td + routes[i].Legs[j].Steps[k].Duration
			}
		}
		tds = append(tds, td)
	}
	fmt.Println(tds)

}
