package main

import (
	"database/sql"
	"errors"

	"./db"
	"./downloader"
	"./parser"
	// "github.com/likipiki/goSerialLoader/db"
	// "github.com/likipiki/goSerialLoader/downloader"
	// "github.com/likipiki/goSerialLoader/parser"
)

var (
	DB       *sql.DB
	pathTest = "./testFiles/" // test directory
	path     = ""
	uid      = "5972916"
	usess    = "74a1217e8a3a5d304d47353cb9db7d57"
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

			resolutionInt, err := getIntResolution(resolution)
			if err != nil {
				return nil, err
			}

			serialsToDownload = append(
				serialsToDownload,
				SerialToDownload{
					Link:     serial.Resolutions[resolutionInt].Link,
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
	err := errors.New("resolution not \"SD\", \"1080p\", \"MP4\"")
	return -1, err
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
