package service

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repository"
)

type HubService struct {
	repo *repository.Repository
}

func NewHub(repo *repository.Repository) *HubService {
	return &HubService{
		repo: repo,
	}
}
func (s *HubService) AllServers() ([]models.Room, error) {
	return s.repo.AllServers()
}

func (s *HubService) GetServerById(id_room int) (models.Room, error) {
	return s.repo.GetServerById(id_room)
}

func (s *HubService) CreateServer(username string) (int, error) {
	return s.repo.CreateServer(username)
}

func (s *HubService) Connect(room_id int, username string) (models.Room, error) {
	return s.repo.Connect(room_id, username)
}

func (s *HubService) Disconnect(room_id int, username string) (models.Room, error) {
	return s.repo.Disconnect(room_id, username)
}

func (s *HubService) IsServerFull(id_room int) (bool, error) {
	return s.repo.IsServerFull(id_room)
}

func (s *HubService) GetUsernameById(id int) (string, error) {
	return s.repo.GetUsernameById(id)
}

func (s *HubService) GetUsersByRoomId(id int) []int {
	return s.repo.GetUsersByRoomId(id)
}
