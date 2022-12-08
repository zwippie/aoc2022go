package myinput

import (
	"embed"
	"fmt"
	"log"
)

//go:embed data/*.txt
var dataSets embed.FS

func ReadInput(day string, suffix string) []byte {
	data, err := dataSets.ReadFile(fmt.Sprintf("data/day%s%s.txt", day, suffix))
	if err != nil {
		log.Fatal(err)
	}
	return data
}
