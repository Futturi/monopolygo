package repository

import (
	"awesomeProject/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type HubPostgres struct {
	db *sqlx.DB
}

func NewHubPostgres(db *sqlx.DB) *HubPostgres {
	return &HubPostgres{db: db}
}

func (r *HubPostgres) AllServers() ([]models.Room, error) {
	var rooms []models.Room
	query := fmt.Sprintf("SELECT room_id,first_player_id, second_player_id, third_player_id,fourth_player_id FROM %s", roomTable)
	err := r.db.Select(&rooms, query)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *HubPostgres) GetServerById(id_room int) (models.Room, error) {
	var room models.Room
	query := fmt.Sprintf("SELECT room_id, first_player_id, second_player_id, third_player_id,fourth_player_id FROM %s WHERE room_id = $1", roomTable)
	err := r.db.Get(&room, query, id_room)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func (r *HubPostgres) CreateServer(username string) (int, error) {
	var room_id int
	user_id, err := r.GetUserIdByUsername(username)
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s(first_player_id) values($1) RETURNING room_id", roomTable)
	row := r.db.QueryRow(query, user_id)
	err = row.Scan(&room_id)
	if err != nil {
		return 0, err
	}
	return room_id, nil
}

func (r *HubPostgres) GetUserIdByUsername(username string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE username = $1", userTale)
	row := r.db.QueryRow(query, username)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
