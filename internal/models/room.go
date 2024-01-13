package models

type Room struct {
	RoomId           int `json:"room_id" db:"room_id"`
	First_player_id  int `json:"first_player_id" db:"first_player_id"`
	Second_player_id int `json:"second_player_id" db:"second_player_id"`
	Third_player_id  int `json:"third_player_id" db:"third_player_id"`
	Fourth_player_id int `json:"fourth_player_id" db:"fourth_player_id"`
}
