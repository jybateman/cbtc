package main

import (
	"os"
	"fmt"
	"errors"
	"io/ioutil"
	"path/filepath"
		
	"github.com/jybateman/gobencode"
)

type torrentInfo struct {
	Name string
	Announce string
	Length int
	PieceLength int
	PieceHash [][]byte
	fileDesc *os.File
	// MetaInfo map[string]interface{}
}

func (t *torrentInfo) getPiecesHash(pieces string) {
	for i := 0; i < len(pieces); i += 20 {
		t.PieceHash = append(t.PieceHash, []byte(pieces)[i:i+20])
	}
}

// TODO
// Make this an array of string
// Better error management
func newTorrent() torrentInfo {
	var torrent torrentInfo
	buf, err := ioutil.ReadFile(fileName)
	checkError(err)
	metaInfo, err := bencode.Decode(string(buf))
	checkError(err)
	if (!isType("string", metaInfo["announce"])) {
		checkError(errors.New("Invalid announce in torrent file"))
	}
	torrent.Announce = metaInfo["announce"].(string)
	
	if (!isType("map[string]interface {}", metaInfo["info"])) {
		checkError(errors.New("Invalid torrent metainfo"))
	}
	
	info := metaInfo["info"].(map[string]interface{})
	if !isType("string", info["name"]) {
		_, f := filepath.Split(fileName)
		torrent.Name = f
	} else {
		torrent.Name = info["name"].(string)
	}

	if !isType("int", info["piece length"]) {
		checkError(errors.New("Invalid piece length in torrent file"))
	} 
	torrent.PieceLength = info["piece length"].(int)

	if !isType("int", info["length"]) {
		checkError(errors.New("Invalid length in torrent file"))
	} 
	torrent.Length = info["length"].(int)

	if !isType("string", info["pieces"]) || len(info["pieces"].(string)) % 20 != 0 {
		checkError(errors.New("Invalid peices in torrent file"))
	}
	torrent.getPiecesHash(info["pieces"].(string))
	fmt.Println(len(torrent.PieceHash))
	return torrent
}
