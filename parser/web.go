package parser

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"../db"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Description   string   `xml:"description"`
	Link          string   `xml:"link"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Language      string   `xml:"language"`
	Items         []Item   `xml:"item"`
}

type Item struct {
	XMLName  xml.Name `xml:"item"`
	Title    string   `xml:"title"`
	Category string   `xml:"category"`
	PubDate  string   `xml:"pubDate"`
	Link     string   `xml:"link"`
}

type Serial struct {
	db.Serial
	Resolutions []Resolution
}

type Resolution struct {
	Format string
	Link   string
}

func Parse(file string) ([]Serial, error) {
	var rss Rss
	err := xml.Unmarshal([]byte(file), &rss)
	if err != nil {
		return nil, err
	}
	var serials []Serial
	for i := 0; i < 5; i++ {
		var resolutions []Resolution
		var serialDb db.Serial
		strs := strings.Split(rss.Channel.Items[i*3].Title, "(")
		serialDb.Name = strings.Split(strs[1], ")")[0]
		serialDb.SeasonData = strings.Split(string(strs[2]), ")")[0]
		for j := 0; j < 3; j++ {
			resolutions = append(resolutions, Resolution{
				strings.Split(strings.Split(rss.Channel.Items[i*3+j].Title, "[")[1], "]")[0],
				rss.Channel.Items[i*3+j].Link})
		}
		serial := Serial{serialDb, resolutions}
		serials = append(serials, serial)
	}
	return serials, nil
}

func Download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
