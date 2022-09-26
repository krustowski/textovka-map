package main

import (
	"math"
	"math/rand"
)

func (m *Map) Init(name string, roomCount int8) {
	m.MapMeta.Name = name
	m.MapMeta.RoomCount = roomCount

	// Generate only 2D square maps!
	m.MapMeta.Dimensions = 2
}

func (m *Map) GenerateRooms() {
	var (
		j           int8
		roomNameLen int = 8
		roomNames       = []string{}
	)

	for j = 0; j < int8(m.MapMeta.RoomCount); j++ {
		name := RandStringBytes(roomNameLen)
		roomNames = append(roomNames, name)
	}

	// Init rooms map.
	rooms := make(map[string]Room)

	// Iterate over room names, allocate rooms.
	for _, room := range roomNames {
		rooms[room] = Room{}
	}

	// Assign room names and rooms themselves to the map.
	m.Rooms = rooms
	m.RoomNames = roomNames
}

// GenerateMapField can only generate 2D square maps atm...
func (m *Map) GenerateMapField() {
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
	//var premap [i][i]int
}

func (m *Map) LinkRooms() {
	directions := []string{"north", "south", "west", "east"}

	// random number -- room index -- nope, start from the first room name
	for _, room := range m.RoomNames {
		var i int

		// random [0 <= n < 4] rooms to assign
		linkCount := rand.Intn(4)

		// loop over links, assign new links to the room
		for i = 0; i < linkCount; i++ {

			// random direction
			direction := directions[i]

			// assign random room from rooms count number
			roomCount := int(m.MapMeta.RoomCount)
			randomRoomIdx := rand.Intn(roomCount)
			randomRoom := m.RoomNames[randomRoomIdx]

			// https://stackoverflow.com/a/69006398
			// First we get a "copy" of the entry
			if entry, ok := m.Rooms[room]; ok {

				// Then we modify the copy
				switch direction {
				case "north":
					entry.North = randomRoom
				case "south":
					entry.South = randomRoom
				case "west":
					entry.West = randomRoom
				case "east":
					entry.East = randomRoom
				}

				// Then we reassign map entry
				m.Rooms[room] = entry
			}
		}
	}
}
