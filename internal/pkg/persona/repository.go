package persona

import "RSOI/internal/models"

type IRepository interface {
	Insert(persona *models.Persona) (uint, error)
	Select(id uint) (*models.Persona, error)
	Update(id uint, persona *models.Persona) error
	Delete(id uint) error
}
