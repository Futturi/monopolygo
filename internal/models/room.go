package models

type Room struct {
	RoomId   int `json:"room_id" db:"room_id"`
	MaxUsers int `json:"maxusers" db:"maxusers"`
}
