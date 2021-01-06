package main

import (
	"os"
	"flag"
	"io/ioutil"
	"encoding/json"
)

var fileName string

var torrents []torrentInfo

func saveData() {
	data, err := json.Marshal(torrents)
	checkError(err)
	f, err := os.OpenFile(".torrentData", os.O_RDWR|os.O_CREATE, 0644)
	checkError(err)
	defer f.Close()
	err = f.Truncate(0)
	checkError(err)
	_, err = f.Write(data)
	checkError(err)
}

func loadData() {
	data, err := ioutil.ReadFile(".torrentData")
	checkError(err)
	err = json.Unmarshal(data, &torrents)
	checkError(err)
}

func initFlag() {
	flag.StringVar(&fileName, "f", "", "File name")
	flag.Parse()
}

func main() {
	loadData()
	initFlag()
	if (fileName != "") {
		torrents = append(torrents, newTorrent())
	}
	saveData()
}
