package main

import (
	"fmt"
	"testing"
)

func TestGetIntResolution(t *testing.T) {
	resolution, err := getIntResolution("MP4")
	if err != nil {
		t.Error("error in function\n" + err.Error())
	}
	if resolution != 0 {
		t.Error("expected 0, got " + string(resolution))
	}
	resolution, err = getIntResolution("1080p")
	if err != nil {
		t.Error("error in function\n" + err.Error())
	}
	if resolution != 1 {
		t.Error("expected 1, got " + string(resolution))
	}
	resolution, err = getIntResolution("SD")
	if err != nil {
		t.Error("error in function\n" + err.Error())
	}
	if resolution != 2 {
		t.Error("expected 2, got " + string(resolution))
	}
}

func TestMainWithOutDB(t *testing.T) {
	serials, err := parseSerials()
	if err != nil {
		t.Error("error in parseSerials() function\n" + err.Error())
	}
	var serialsToDownload []SerialToDownload
	for i, serial := range serials {

		serialsToDownload = append(
			serialsToDownload,
			SerialToDownload{
				Link:     serial.Resolutions[1].Link,
				FileName: fmt.Sprint(i, ".torrent"),
			},
		)

	}
	err = downloadSerials(serialsToDownload, uid, usess, pathTest)
	if err != nil {
		t.Error("error in downloadSerials() function\n" + err.Error())
	}
}
