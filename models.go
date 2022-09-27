package main

type Map struct {
	// Additional meta map pamameters, passed as flags to generator.
	MapMeta MapMeta `json:"meta"`

	// Name of the initial game room.
	StartRoom string `json:"start_room"`

	// End-of-the-game room name.
	EndRoom string `json:"end_room"`

	// List of rooms.
	RoomNames []string `json:"rooms"`

	// Map of rooms.
	Rooms map[string]Room `json:"room"`
}

type MapMeta struct {
	// Generated map name.
	Name string `json:"map_name"`

	// Map dimensions (only 2 dimensions are implemented for now).
	Dimensions int8 `json:"map_dimensions"`

	// Number of rooms for such map.
	RoomCount int8 `json:"room_count"`
}

type Room struct {
	// Shown descriptior of such room (e.g. Dark room with a mysterious locked door to west).
	Description string `json:"description"`

	// Name of the room linked to north.
	North string `json:"north"`

	// Name of the room linked to south.
	South string `json:"south"`

	// Name of the room linked to west.
	West string `json:"west"`

	// Name of the room linked to east.
	East string `json:"east"`

	// List of items to grab to inventory.
	Items []string `json:"items"`

	// Named list of objects in room.
	Objects []string `json:"objects"`

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
