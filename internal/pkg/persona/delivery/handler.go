package delivery

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
	"strconv"
)

type PHandler struct {
	personaUsecase persona.IUsecase
}

func NewPHandler(personaUsecase persona.IUsecase) *PHandler {
	return &PHandler{personaUsecase: personaUsecase}
}

func (h *PHandler) Create(w http.ResponseWriter, r *http.Request) {

	person := &models.PersonaRequest{}
	err := easyjson.UnmarshalFromReader(r.Body, person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	v := validator.New()
	errs := v.Struct(person)
	if errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			log.Println(e)
		}
	}

	id, code := h.personaUsecase.Create(person)

	switch code {
	case models.OKEY:
		log.Print(id)
		w.WriteHeader(http.StatusCreated)
	case models.NOTFOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *PHandler) Read(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	ids := v["personID"]
	id, _ := strconv.Atoi(ids)
	_, code := h.personaUsecase.Read(uint(id))

	switch code {
	case models.OKEY:
		log.Print(id)
		w.WriteHeader(http.StatusOK)
	case models.NOTFOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (h *PHandler) ReadAll(w http.ResponseWriter, r *http.Request) {

	_, code := h.personaUsecase.ReadAll()
	switch code {
	case models.OKEY:
		w.WriteHeader(http.StatusOK)
	case models.NOTFOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *PHandler) Update(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	ids := v["personID"]
	id, _ := strconv.Atoi(ids)

	persona := &models.PersonaRequest{}
	easyjson.UnmarshalFromReader(r.Body, persona)

	code := h.personaUsecase.Update(uint(id), persona)

	switch code {
	case models.OKEY:
		w.WriteHeader(http.StatusCreated)
	case models.NOTFOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func (h *PHandler) Delete(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	ids := v["personID"]
	id, _ := strconv.Atoi(ids)
	code := h.personaUsecase.Delete(uint(id))

	switch code {
	case models.OKEY:
		w.WriteHeader(http.StatusCreated)
	case models.NOTFOUND:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
