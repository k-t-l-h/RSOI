package persona

import "RSOI/internal/models"

type IUsecase interface {
	Create(persona *models.PersonaRequest) (uint, error)
	Read(id uint) (*models.PersonaResponse, error)
	ReadAll() ([]*models.PersonaResponse, error)
	Update(id uint, persona *models.PersonaRequest) error
	Delete(id uint) error
}
