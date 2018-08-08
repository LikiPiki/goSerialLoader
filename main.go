package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/likipiki/goSerialLoader/db"
	"github.com/likipiki/goSerialLoader/downloader"
	"github.com/likipiki/goSerialLoader/parser"
)

// main BD
var (
	DB       *sql.DB
	pathTest = "./testsFiles/"            // test directory
	path     = "/Users/sergey/Downloads/" // example
	uid      = "5972916"
	usess    = "74a1217e8a3a5d304d47353cb9db7d57"
)

// link to lostfilm rss
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
	err = downloadSerials(serialsToDownload, uid, usess, path)
	if err != nil {
		panic(err)
	}
}

func parseSerials() ([]parser.Serial, error) {
	file, err := parser.Download(LINK)
	if err != nil {
		return nil, err
	}
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
		fmt.Println("serial is ", serial)
		oldSerial, err := serial.Get()
		if err != nil {
			return nil, err
		}
		fmt.Println("serial is", oldSerial, db.Serial{})
		if oldSerial == (db.Serial{}) {
			continue
		}

		if serial.Serial.Season > oldSerial.Season || serial.Serial.Episode > oldSerial.Episode {

			err := serial.UpdateSeasonEpisode()
			if err != nil {
				return nil, err
			}

			resolitionInt, err := getIntResolution(oldSerial.Resolution)
			if err != nil {
				return nil, err
			}

			serialsToDownload = append(
				serialsToDownload,
				SerialToDownload{
					Link:     serial.Resolutions[resolitionInt].Link,
					FileName: serial.Serial.Name + " " + serial.SeasonData + ".torrent",
				},
			)

		}

	}
	return serialsToDownload, nil
}

func getIntResolution(resolution string) (int, error) {
	switch resolution {
	case "MP4":
		return 0, nil
	case "1080p":
		return 1, nil
	case "SD":
		return 2, nil
	}
	return -1, errors.New("resolution not \"SD\", \"1080p\", \"MP4\"")
}

func downloadSerials(serials []SerialToDownload, uid, usess, filepath string) error {
	downloader := downloader.Downloader{
		Uid:      uid,
		Usess:    usess,
		Filepath: filepath,
	}
	for _, serial := range serials {
		err := downloader.DownloadTorrentFile(serial.Link, serial.FileName)
		if err != nil {
			return err
		}
	}
	return nil
}
