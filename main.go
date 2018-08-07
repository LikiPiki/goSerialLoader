package main

import (
	"fmt"

	"github.com/likipiki/goSerialLoader/parser"
)

const (
	LINK = "http://retre.org/rssdd.xml"
)

func main() {
	file, err := parser.Download(LINK)
	if err != nil {
		return
	}
	serials, err := parser.Parse(file)
	if err != nil {
		return
	}
	fmt.Println(serials)
}
