package repository

import (
	"RSOI/internal/models"
	"context"
	"github.com/jackc/pgx/pgxpool"
)

type PRepository struct {
	pool pgxpool.Pool
}

func NewPRepository() *PRepository {
	return &PRepository{}
}

func (pr *PRepository) Insert(persona *models.PersonaRequest) (uint, int) {
	res := pr.pool.QueryRow(context.Background(), CREATEPERSONA)
	err := res.Scan(persona.ID)
	if err != nil {
		return 0, models.NOTFOUND
	} else {
		return persona.ID, models.OKEY
	}
}

func (pr *PRepository) Select(id uint) (*models.PersonaResponse, int) {
	res := pr.pool.QueryRow(context.Background(), READPERSONA)
	persona := models.PersonaResponse{ID:id}

	err := res.Scan(persona.Name, persona.Address, persona.Age, persona.Work)
	if err != nil {
		return &persona, models.NOTFOUND
	} else {
		return &persona, models.OKEY
	}
}

func (pr *PRepository) SelectAll() ([]*models.PersonaResponse, int) {
	tag, err := pr.pool.Query(context.Background(), READALLPERSONA)
	var personas []*models.PersonaResponse

	for tag.Next() {
		persona := models.PersonaResponse{}
		err = tag.Scan(persona.Name, persona.Address, persona.Age, persona.Work)
		if err != nil {
			break
		}
		personas = append(personas, &persona)
	}

	if err != nil {
		return personas, models.NOTFOUND
	} else {
		return personas, models.OKEY
	}
}

func (pr *PRepository) Update(id uint, persona *models.PersonaRequest) int {
	tag, err := pr.pool.Exec(context.Background(), UPDATEPERSONA, id)
	if err != nil {
		return models.BADREQUEST
	}
	if tag.RowsAffected() == 0 {
		return models.NOTFOUND
	}

	return models.OKEY


}

func (pr *PRepository) Delete(id uint) int {
	tag, err := pr.pool.Exec(context.Background(), UPDATEPERSONA, id)
	if err != nil {
		return models.BADREQUEST
	}
	if tag.RowsAffected() == 0 {
		return models.NOTFOUND
	}

	return models.OKEY
}
