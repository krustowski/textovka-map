package main

type Map struct {
	MapMeta   MapMeta `json:"meta"`
	StartRoom string  `json:"start_room"`
	EndRoom   string  `json:"end_room"`

	// List of rooms.
	RoomNames []string `json:"rooms"`

	// Map of rooms.
	Rooms map[string]Room `json:"room"`
}

type MapMeta struct {
	Name       string `json:"map_name"`
	Dimensions int8   `json:"map_dimensions"`
	RoomCount  int8   `json:"room_count"`
}

type Room struct {
	// Shown descriptior of such room (e.g. Dark room with a mysterious locked door to west).
	Description string   `json:"description"`
	North       string   `json:"north"`
	South       string   `json:"south"`
	West        string   `json:"west"`
	East        string   `json:"east"`
	Items       []string `json:"items"`
	Objects     []string `json:"objects"`

	// List of actions -- linked to their effects.
	Actions []string `json:"actions"`

	// Map of actions' effects.
	Effects map[string]Effect `json:"effects"`
}

type Effect struct {
	Type         string `json:"type"`
	RequiredItem string `json:"required-item"`
	ObjectTarget string `json:"object"`
	ShowHidden   bool   `json:"show-hidden"`
	DamageHP     []int  `json:"damage-hp"`
}
