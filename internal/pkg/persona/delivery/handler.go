package delivery

import (
	"RSOI/internal/pkg/persona"
	"net/http"
)

type PHandler struct {
	personaUsecase persona.IUsecase
}

func (h *PHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *PHandler) Read(w http.ResponseWriter, r *http.Request) {

}

func (h *PHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *PHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
