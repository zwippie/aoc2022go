// blabla
package myinput

import (
	"embed"
	"fmt"
	"log"
)

//go:embed data/*
var dataSets embed.FS

// Read a file from the data directory.
func ReadInput(fileName string) []byte {
	data, err := dataSets.ReadFile(fmt.Sprintf("data/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	return data
}
