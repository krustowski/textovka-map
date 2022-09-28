package main

import "flag"

var (
	flagMapName   *string
	flagRoomCount *int
)

func parseFlags() {
	flagMapName = flag.String("mapName", "testMap", "a string, name of the new map")
	flagRoomCount = flag.Int("roomCount", 10, "an int, number of rooms to be generated")

	flag.Parse()
}
