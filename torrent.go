package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"path/filepath"
		
	"github.com/jybateman/gobencode"
)

type torrentInfo struct {
	Name string
	Announce string
	PieceLength int
	PieceHash [20][]byte
	MetaInfo map[string]interface{}
}

// TODO
// Make this an array of string
// Better error management
func newTorrent() torrentInfo {
	var torrent torrentInfo
	buf, err := ioutil.ReadFile(fileName)
	checkError(err)
	torrent.MetaInfo, err = bencode.Decode(string(buf))
	checkError(err)
	if (!isType("map[string]interface {}", torrent.MetaInfo["info"])) {
		checkError(errors.New("Invalid torrent metainfo"))
	}
	info := torrent.MetaInfo["info"].(map[string]interface{})
	fmt.Println(info["name"])
	if (!isType("string", info["name"])) {
		_, f := filepath.Split(fileName)
		torrent.Name = f
	} else {
		torrent.Name = info["name"].(string)
	}
	return torrent
}
