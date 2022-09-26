package main

import "flag"

var (
	MapName   *string
	RoomCount *int
)

func parseFlags() {
	MapName = flag.String("mapName", "testMap", "a string, name of the new map")
	RoomCount = flag.Int("roomCount", 10, "an int, number of rooms to be generated")

	flag.Parse()
}
