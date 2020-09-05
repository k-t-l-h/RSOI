package delivery

import (
	"RSOI/internal/pkg/persona"
	"net/http"
)

type PHandler struct {
	personaUsecase persona.IUsecase
}

func NewPHandler(personaUsecase persona.IUsecase) *PHandler {
	return &PHandler{personaUsecase: personaUsecase}
}

func (h *PHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (h *PHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *PHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *PHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *PHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
