package persona

import "RSOI/internal/models"

type IRepository interface {
	Insert(persona *models.PersonaRequest) (uint, error)
	Select(id uint) (*models.PersonaResponse, error)
	SelectAll() ([]*models.PersonaResponse, error)
	Update(id uint, persona *models.PersonaRequest) error
	Delete(id uint) error
}
