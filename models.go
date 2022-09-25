package main

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
