package delivery

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
	"RSOI/src/github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"log"
	"net/http"
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

	id, err := h.personaUsecase.Create(person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		//TODO: добавить обработку id
		log.Print(id)
		w.WriteHeader(http.StatusCreated)
	}
}

func (h *PHandler) Read(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["personID"]
	_, err := h.personaUsecase.Read(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func (h *PHandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *PHandler) Update(w http.ResponseWriter, r *http.Request) {

	v := mux.Vars(r)
	id := v["personID"]
	err := h.personaUsecase.Update(id, nil)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func (h *PHandler) Delete(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["personID"]
	err := h.personaUsecase.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
