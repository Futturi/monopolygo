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
