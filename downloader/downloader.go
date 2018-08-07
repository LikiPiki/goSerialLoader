package downloader

import (
	"io"
	"net/http"
	"os"
)

type Downloader struct {
	Uid   string
	Usess string
}

func (downloader Downloader) DownloadTorrentFile(link string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", "uid="+downloader.Uid+";usess="+downloader.Usess)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
