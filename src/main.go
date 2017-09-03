package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Entry struct {
	Title string `xml:"title"`
}

type Entrys struct {
	Entrys []Entry `xml:"entry"`
}

func main() {

	response, err := http.Get("https://alerts.weather.gov/cap/us.php?x=0")
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting XML feed from NOAA:", err)
		return
	}

	var q Entrys
	xml.Unmarshal(body, &q)
	fmt.Println(q)
}
