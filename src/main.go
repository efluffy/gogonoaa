package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Geocode struct {
	ValueName string `xml:"valueName"`
	Value     string `xml:"value"`
}

type Entry struct {
	Title     string  `xml:"title"`
	Updated   string  `xml:"updated"`
	Published string  `xml:"published"`
	Summary   string  `xml:"summary"`
	AlertType string  `xml:"msgType"`
	Location  Geocode `xml:"geocode"`
}

type Entrys struct {
	Entrys []Entry `xml:"entry"`
}

func main() {
	var capcode = os.Args[1]
	response, err := http.Get("https://alerts.weather.gov/cap/us.php?x=0")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting XML feed from NOAA:", err)
		return
	}

	var q Entrys
	xml.Unmarshal(body, &q)

	for _, Entry := range q.Entrys {
		if Entry.Location.Value == capcode {
			fmt.Printf("%s : %s\n", Entry.AlertType, Entry.Title)
			fmt.Printf("%s", Entry.Published)
			fmt.Printf(" - %s\n", Entry.Updated)
			fmt.Printf("\t%s\n", Entry.Summary)
		}
	}
}
