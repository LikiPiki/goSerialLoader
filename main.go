package main

import (
	"fmt"

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

	fmt.Printf("%+v", serials[0])

}
