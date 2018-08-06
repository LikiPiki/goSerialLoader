package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Parse(file string) error {
	var rss Rss
	err := xml.Unmarshal([]byte(file), &rss)
	if err != nil {
		return err
	}
	fmt.Println(rss.Channel.Items[0].Title)
	fmt.Println(rss.Channel.Items[1].Title)
	fmt.Println(rss.Channel.Items[2].Title)

	return nil
}

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
