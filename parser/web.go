package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/likipiki/goSerialLoader/db"
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
		rSerial := regexp.MustCompile(`(S\d\dE\d\d)`)
		rTitle := regexp.MustCompile(`(\((.*)\). )`)

		title := rss.Channel.Items[i*3].Title
		serialDb := db.Serial{
			Name:       rTitle.FindStringSubmatch(title)[2],
			SeasonData: rSerial.FindStringSubmatch(title)[0],
		}

		var resolutions []Resolution
		rResolution := regexp.MustCompile(`\[(\w+)\]`)
		for j := 0; j < 3; j++ {
			current := rss.Channel.Items[i*3+j]
			resolutions = append(resolutions, Resolution{
				rResolution.FindStringSubmatch(current.Category)[1],
				current.Link,
			})
		}

		serials = append(serials, Serial{
			serialDb, resolutions,
		})
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

func DownloadTorrentFile(link string, filepath string, uid string, usess string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", "uid="+uid+";usess="+usess)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.StatusCode)
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
