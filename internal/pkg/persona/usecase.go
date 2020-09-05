package persona

import "RSOI/internal/models"

type IUsecase interface {
	Create(persona *models.Persona) (uint, error)
	Read(id uint) (*models.Persona, error)
	ReadAll() ([]*models.Persona, error)
	Update(id uint, persona *models.Persona) error
	Delete(id uint) error
}
