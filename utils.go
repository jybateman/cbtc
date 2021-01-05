package main

import (
	"fmt"
	"log"
)

// TODO
// replace this with actual error handling
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Check if an interface is of a specified type
// wType is the type we want to conpare the interface with
// inter if th interface we wish to check
func isType(wType string, inter interface{}) bool {
	return wType == fmt.Sprintf("%T", inter)
}
