package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var (
		m         Map    = Map{}
		mapName   string = "testMap"
		roomCount int8   = 7
	)

	m.Init(mapName, roomCount)
	m.GenerateRooms()
	m.GenerateMapField()
	m.LinkRooms()

	// Marshal Map struct into intented JSON.
	j, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Println(err)
		log.Panic("Map struct to JSON marshalling failed...")
	}

	// Write JSON to stdout.
	os.Stdout.Write(j)
}
