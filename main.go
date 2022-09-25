package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"os"
)

type Map struct {
	MapMeta   MapMeta `json:"meta"`
	StartRoom string  `json:"start_room"`
	EndRoom   string  `json:"end_room"`

	// List of rooms.
	RoomNames []string `json:"rooms"`

	// Map of rooms.
	//Rooms interface{} `json:"room"`
	Rooms map[string]Room `json:"room"`
}

type MapMeta struct {
	Name       string `json:"map_name"`
	Dimensions int8   `json:"map_dimensions" default:'2'`
	RoomCount  int8   `json:"room_count"`
}

type Room struct {
	Description string   `json:"description"`
	North       string   `json:"north"`
	South       string   `json:south`
	West        string   `json:"west"`
	East        string   `json:"east"`
	Items       []string `json:"items"`
	Objects     []string `json:"objects"`

	// List of effects.
	Actions []string `json:"actions"`

	// Map of effects.
	Effects interface{} `json:"effects"`
}

type Effect struct {
	Type         string `json:"type"`
	RequiredItem string `json:"required-item"`
	ShowHidden   bool   `json:"show-hidden" default:false`
}

//
//
//

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (m *Map) Init(name string, roomCount int8) {
	m.MapMeta.Name = name
	m.MapMeta.RoomCount = roomCount

	// Generate only 2D square maps!
	m.MapMeta.Dimensions = 2
}

func (m *Map) GenerateRoomNames() {
	var (
		j           int8
		roomNameLen int = 8
		roomNames       = []string{}
	)

	for j = 0; j < int8(m.MapMeta.RoomCount); j++ {
		name := RandStringBytes(roomNameLen)
		roomNames = append(roomNames, name)
	}

	m.RoomNames = roomNames
}

// GenerateMap can only generate 2D square maps atm...
func (m *Map) GenerateMap() {
	var (
		i          float64
		dimensions = float64(m.MapMeta.Dimensions)
		roomCount  = float64(m.MapMeta.RoomCount)
		//premap     interface{}
	)

	// Generate map dimension (e.g. x, y square system) from Dimensions count
	for i = 0; roomCount > math.Pow(dimensions, i); i++ {
	}
	// Output here is the x = y = i square dimensions...

	roomNames := m.RoomNames

	// Init rooms map.
	rooms := make(map[string]Room)

	// Iterate over room names, allocate rooms.
	for _, room := range roomNames {
		rooms[room] = Room{}
	}

	// Assign room names and rooms themselves to the map.
	m.RoomNames = roomNames
	m.Rooms = rooms
}

func main() {
	var m Map = Map{}

	m.Init("testMap", 10)
	m.GenerateRoomNames()
	m.GenerateMap()

	// Marshal Map struct into intented JSON.
	j, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Panic("idk")
	}

	// Write JSON to stdout.
	os.Stdout.Write(j)
}
