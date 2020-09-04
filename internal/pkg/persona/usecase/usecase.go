package usecase

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
)

type PUsecase struct {
	repo persona.IRepository
}

func (us *PUsecase) Create(persona *models.Persona) (uint, error) {
	return 0, nil
}

func (us *PUsecase) Read(id uint) (*models.Persona, error) {
	return nil, nil
}

func (us *PUsecase) Update(id uint, persona *models.Persona) error {
	return nil
}

func (us *PUsecase) Delete(id uint) error {
	return nil
}
