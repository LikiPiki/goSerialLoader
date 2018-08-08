package main

import (
	"database/sql"
	"fmt"

	"./db"
	"./downloader"
	"./parser"
	// "github.com/likipiki/goSerialLoader/db"
	// "github.com/likipiki/goSerialLoader/downloader"
	// "github.com/likipiki/goSerialLoader/parser"
)

var (
	DB    *sql.DB
	path  = "" // /Users/ilja/Downloads/
	uid   = ""
	usess = ""
)

const (
	LINK = "http://retre.org/rssdd.xml"
)

type SerialToDownload struct {
	Link     string
	FileName string
}

func main() {
	DB = db.Connect()
	defer DB.Close()

	serials, err := parseSerials()
	if err != nil {
		panic(err)
	}
	serialsToDownload, err := checkSerials(serials)
	if err != nil {
		panic(err)
	}
	err = downloadSerials(serialsToDownload, uid, usess)
	if err != nil {
		panic(err)
	}
}

func parseSerials() ([]parser.Serial, error) {
	file, err := parser.Download(LINK)
	if err != nil {
		return nil, err
	}
	fmt.Println(file)
	serials, err := parser.Parse(file)
	if err != nil {
		return nil, err
	}
	return serials, nil
}

func checkSerials(serials []parser.Serial) ([]SerialToDownload, error) {
	// fileName contains resolution (SirenaS01E01.torrent)
	var serialsToDownload []SerialToDownload

	for _, serial := range serials {

		oldSeason, oldEpisode, err := serial.Get()
		if err != nil {
			return nil, err
		}

		if serial.Serial.Season > oldSeason || serial.Serial.Episode > oldEpisode {

			err := serial.Set()
			if err != nil {
				return nil, err
			}

			resolution, err := serial.GetResolution()
			if err != nil {
				return nil, err
			}

			var resulitionInt int

			switch resolution {
			case "MP4":
				resulitionInt = 0
			case "1080p":
				resulitionInt = 1
			case "SD":
				resulitionInt = 2
			}

			serialsToDownload = append(
				serialsToDownload,
				SerialToDownload{
					Link:     serial.Resolutions[resulitionInt].Link,
					FileName: serial.Serial.Name + " " + serial.SeasonData + ".torrent",
				},
			)

		}

	}
	return serialsToDownload, nil
}

func downloadSerials(serials []SerialToDownload, uid, usess string) error {
	downloader := downloader.Downloader{
		Uid:   uid,
		Usess: usess,
	}
	for _, serial := range serials {
		err := downloader.DownloadTorrentFile(serial.Link, path+serial.FileName)
		if err != nil {
			return err
		}
	}
	return nil
}
