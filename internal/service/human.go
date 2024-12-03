package service

import (
	"passport/internal/models"
	"passport/internal/repository"
	"passport/internal/validators"
)

type HumanService struct {
	repo *repository.HumanRepository
}

func NewHumanService(repo *repository.HumanRepository) *HumanService {
	return &HumanService{repo: repo}
}

func (s *HumanService) CreateHuman(human *models.Human) error {
	if err := validators.ValidateHuman(human); err != nil {
		return err
	}

	return s.repo.CreateHuman(human)
}

func (s *HumanService) GetHuman(human *models.Human) (*models.Human, error) {
	if err := validators.ValidateHuman(human); err != nil {
		return nil, err
	}

	return s.repo.GetHuman(human.ID)
}

func (s *HumanService) UpdateHuman(human *models.Human) error {
	if err := validators.ValidateHuman(human); err != nil {
		return err
	}

	return s.repo.UpdateHuman(human)
}

func (s *HumanService) DeleteHuman(human *models.Human) error {
	if err := validators.ValidateHuman(human); err != nil {
		return err
	}

	return s.repo.DeleteHuman(human)
}

func (s *HumanService) GetAllHumans() ([]*models.Human, error) {
	return s.repo.GetAllHumans()
}

func (s *HumanService) GetHumanBySerialNumber(serialNumber string) (*models.Human, error) {
	return s.repo.GetHumanBySerialNumber(serialNumber)
}
