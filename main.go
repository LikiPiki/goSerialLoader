package main

import (
	"database/sql"
	"fmt"

	"github.com/likipiki/goSerialLoader/db"
	"github.com/likipiki/goSerialLoader/parser"
)

var (
	DB *sql.DB
)

const (
	LINK = "http://retre.org/rssdd.xml"
)

func main() {
	DB = db.Connect()
	defer DB.Close()

	one, two, err := db.SerialDB{
		Serial: db.Serial{
			Name: "kek",
		},
	}.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(one, two)

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
