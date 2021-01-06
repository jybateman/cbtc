package main

import (
	"errors"
	"io/ioutil"
		
	"github.com/jybateman/gobencode"
)

type torrentInfo struct {
	Name string
	Announce string
	PieceLength int
	PieceHash [20]byte
	MetaInfo map[string]interface{}
}

func 

// TODO
// Make this an array of string
func addTorrent() torrentInfo {
	var torrent torrentInfo
	buf, err := ioutil.ReadFile(fileName)
	checkError(err)
	torrent.MetaInfo, err = bencode.Decode(string(buf))
	checkError(err)
	if (!isType("map[string]interface {}", torrent.MetaInfo["info"])) {
		checkError(errors.New("Invalid torrent metainfo"))
	}
	return torrent
}
