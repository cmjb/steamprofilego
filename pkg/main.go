package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SteamIdXML struct {
	XMLName   xml.Name `xml:"profile"`
	SteamID64 uint64   `xml:"steamID64"`
	SteamID   string   `xml:"steamID"`
}

func main() {
	arg := os.Args[1]
	err, s := profileToId64(arg)
	if err != nil {
		return
	}
	fmt.Println(s)
}

func profileToId64(url string) (error, uint64) {
	response, err := http.Get(url + "?xml=1")
	if err != nil {
		return err, 0
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err, 0
	}

	var result SteamIdXML
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return err, 0
	}

	return err, result.SteamID64
}
