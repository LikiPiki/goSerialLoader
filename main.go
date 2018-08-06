package main

import (
	"./parser"
)

func main() {
	// err := Parse(download("http://retre.org/rssdd.xml"))
	str, err := parser.Download("http://retre.org/rssdd.xml")
	if err != nil {
		panic(err)
	}

	err = parser.Parse(str)
	if err != nil {
		panic(err)
	}

}
