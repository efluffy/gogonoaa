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
	Alerts []Alert `xml:"info"`
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
	var found = 0
	if len(os.Args) < 2 {
		fmt.Println("Please provide the CAP location code to check.")
		os.Exit(1)
	}
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
			for _, Alert := range qAlert.Alerts {
				found++
				fmt.Printf("<font id=\"alertHeader\">%s</font> - <font id=\"alertBody\">%s</font><br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<font id=\"alertDesc\">%s</font>", Alert.Type, Alert.Headline, Alert.Description)
			}
		}
	}
	if found == 0 {
		fmt.Println("<font id=\"alertBody\">No active watches or warnings for this area.<font>")
	}
	os.Exit(0)
}
