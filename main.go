package main

import (
	"github.com/likipiki/goSerialLoader/parser"
)

func main() {
	str, err := parser.Download("http://retre.org/rssdd.xml")
	if err != nil {
		panic(err)
	}
	var serials []parser.Serial
	serials, err = parser.Parse(str)
	if err != nil {
		panic(err)
	}
	err = parser.DownloadTorrentFile(
		serials[0].Resolutions[2].Link,
		"path to downloads"+serials[0].Serial.Name+serials[0].Serial.SeasonData+".torrent",
		"uid",
		"usess")
	if err != nil {
		panic(err)
	}
}
