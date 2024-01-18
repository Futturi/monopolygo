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
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s(first_player_id) values($1) RETURNING room_id", roomTable)
	row := tx.QueryRow(query, user_id)
	err = row.Scan(&room_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query_rooms_users := fmt.Sprintf("INSERT INTO %s(user_id,room_id) values($1,$2)", usersRooms)

	_, err = tx.Exec(query_rooms_users, user_id, room_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return room_id, tx.Commit()
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

func (r *HubPostgres) Connect(room_id int, username string) (models.Room, error) {
	var room models.Room
	user_id, err := r.GetUserIdByUsername(username)
	if err != nil {
		return models.Room{}, err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return models.Room{}, err
	}

	rooms_query := fmt.Sprintf(`UPDATE %s
		SET
		second_player_id = CASE
			WHEN second_player_id = 1 then $1
			ELSE second_player_id 
			end,
		third_player_id = case
			when second_player_id != 1 and third_player_id = 1 then $1
			else third_player_id 
			end,
		fourth_player_id = case 
			when second_player_id != 1 and third_player_id != 1 and fourth_player_id = 1 then $1
			else fourth_player_id 
			end
		where room_id =$2 returning *;`, roomTable)
	row := tx.QueryRow(rooms_query, user_id, room_id)
	err = row.Scan(&room.RoomId, &room.First_player_id, &room.Second_player_id, &room.Third_player_id, &room.Fourth_player_id)
	if err != nil {
		tx.Rollback()
		return models.Room{}, err
	}

	rooms_user_query := fmt.Sprintf("INSERT INTO %s(user_id,room_id) values($1,$2)", usersRooms)
	_, err = tx.Exec(rooms_user_query, user_id, room_id)
	if err != nil {
		tx.Rollback()
		return models.Room{}, err
	}
	return room, tx.Commit()
}

func (r *HubPostgres) Disconnect(room_id int, username string) (models.Room, error) {
	var room models.Room

	user_id, err := r.GetUserIdByUsername(username)
	if err != nil {
		return models.Room{}, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return models.Room{}, err
	}

	rooms_query := fmt.Sprintf(`UPDATE %s
		SET
		first_plyaer_id = CASE
			WHEN first_player_id = $1 then 1
			ELSE first_player_id
			end,
		second_player_id = CASE
			WHEN second_player_id = $1 then 1
			ELSE second_player_id 
			end,
		third_player_id = case
			when third_player_id = $1 then 1
			else third_player_id 
			end,
		fourth_player_id = case 
			when ourth_player_id = $1 then 1
			else fourth_player_id 
			end
		where room_id =$2 returning *;`, roomTable)
	row := tx.QueryRow(rooms_query, user_id, room_id)

	if err = row.Scan(&room); err != nil {
		return models.Room{}, err
	}

	rooms_user_query := fmt.Sprintf("UPDATE %s SET user_id = 1 WHERE room_id = $1 and user_id = $2", usersRooms)

	_, err = tx.Exec(rooms_user_query, room_id, user_id)
	if err != nil {
		return models.Room{}, err
	}

	return room, tx.Commit()
}

func (r *HubPostgres) IsServerFull(id_room int) (bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE room_id = $1 and first_player_id != 1 and second_player_id != 1 and third_player_id != 1 and fourth_player_id != 1", roomTable)
	_, err := r.db.Exec(query, id_room)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *HubPostgres) GetUsernameById(id int) (string, error) {
	query := fmt.Sprintf("SELECT username FROM %s WHERE user_id = $1", userTale)
	var username string

	err := r.db.QueryRow(query, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (r *HubPostgres) GetUsersByRoomId(id int) []int {

	var users []int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE room_id = $1", usersRooms)
	err := r.db.Select(&users, query, id)
	if err != nil {
		return nil
	}
	return users
}
