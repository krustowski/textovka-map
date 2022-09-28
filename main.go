package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var m Map = Map{}

	// load and parse app flags
	parseFlags()

	// map generation procedure
	m.Init(*flagMapName, *flagRoomCount)
	m.GenerateRooms()
	//m.GenerateMapField()
	m.LinkRooms()

	// next: AddObjects, AddItems, AddActions, AddEffects

	// Marshal Map struct into indented JSON.
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
		log.Panic("Map struct to JSON marshalling failed...")
	}

	// Write JSON to stdout.
	os.Stdout.Write(jsonData)
}
