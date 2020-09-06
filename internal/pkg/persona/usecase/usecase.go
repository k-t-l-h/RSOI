package usecase

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
)

type PUsecase struct {
	repo persona.IRepository
}

func NewPUsecase(repo persona.IRepository) *PUsecase {
	return &PUsecase{repo: repo}
}

func (us *PUsecase) Create(persona *models.PersonaRequest) (uint, error) {
	return 0, nil
}

func (us *PUsecase) Read(id uint) (*models.PersonaResponse, error) {
	return nil, nil
}

func (us *PUsecase) ReadAll() ([]*models.PersonaResponse, error) {
	return nil, nil
}

func (us *PUsecase) Update(id uint, persona *models.PersonaRequest) error {
	return nil
}

func (us *PUsecase) Delete(id uint) error {
	return nil
}
