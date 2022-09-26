package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var (
		m Map = Map{}
	)

	// load and parse app flags
	parseFlags()

	m.Init(*MapName, *RoomCount)
	m.GenerateRooms()
	//m.GenerateMapField()
	m.LinkRooms()

	// next: AddObjects, AddItems, AddActions, AddEffects

	// Marshal Map struct into intented JSON.
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
		log.Panic("Map struct to JSON marshalling failed...")
	}

	// Write JSON to stdout.
	os.Stdout.Write(jsonData)
}
