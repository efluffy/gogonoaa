package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Alert struct {
	Type        string `xml:"event"`
	Issued      string `xml:"effective"`
	Expires     string `xml:"expires"`
	Headline    string `xml:"headline"`
	Description string `xml:"description"`
}

type Alerts struct {
	Alerts []Alert `xml:"alert"`
}

type Geocode struct {
	ValueName string `xml:"valueName"`
	Value     string `xml:"value"`
}

type Entry struct {
	Title    string  `xml:"title"`
	Link     string  `xml:"id"`
	Location Geocode `xml:"geocode"`
}

type Entrys struct {
	Entrys []Entry `xml:"entry"`
}

func main() {
	var capcode = os.Args[1]
	response, err := http.Get("https://alerts.weather.gov/cap/us.php?x=0")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting main XML feed from NOAA:", err)
		return
	}

	var q Entrys
	xml.Unmarshal(body, &q)

	for _, Entry := range q.Entrys {
		re := regexp.MustCompile(capcode)
		if re.MatchString(Entry.Location.Value) {
			responseAlert, err := http.Get(Entry.Link)
			bodyAlert, err := ioutil.ReadAll(responseAlert.Body)
			if err != nil {
				fmt.Println("Error getting alert XML feed from NOAA:", err)
				return
			}
			var qAlert Alerts
			xml.Unmarshal(bodyAlert, &qAlert)
			fmt.Println(len(qAlert.Alerts))
			for _, Alert := range qAlert.Alerts {
				fmt.Printf("%s : %s\n%s - %s\n%s", Alert.Type, Alert.Headline, Alert.Issued, Alert.Expires, Alert.Description)
			}
		}
	}
}
