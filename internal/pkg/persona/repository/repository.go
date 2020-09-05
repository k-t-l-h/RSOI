package repository

import "RSOI/internal/models"

type PRepository struct {
	//здесь база данных
}

func NewPRepository() *PRepository {
	return &PRepository{}
}

func (pr *PRepository) Insert(persona *models.Persona) (uint, error) {
	return 0, nil
}

func (pr *PRepository) Select(id uint) (*models.Persona, error) {
	return nil, nil
}

func (pr *PRepository) SelectAll() ([]*models.Persona, error) {
	return nil, nil
}

func (pr *PRepository) Update(id uint, persona *models.Persona) error {
	return nil
}

func (pr *PRepository) Delete(id uint) error {
	return nil
}
