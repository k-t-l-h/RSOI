package delivery

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
	"RSOI/internal/pkg/persona/mock"
	"fmt"
	"github.com/golang/mock/gomock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPHandler_Create(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock.NewMockIUsecase(ctl)

	type fields struct {
		personaUsecase persona.IUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected models.PersonaRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple create",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/person",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" }`, "name"))),
				expected: models.PersonaRequest{Name: "name"},
				status:   http.StatusCreated,
			}},
		{
			name:   "json err",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r: httptest.NewRequest("POST", "/person",
					strings.NewReader(fmt.Sprintf(`{"name": "%s" `, "name"))),
				expected: models.PersonaRequest{Name: "name"},
				status:   http.StatusBadRequest,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}
			w := httptest.NewRecorder()

			mockUsecase.EXPECT().Create(&tt.args.expected).Return(uint(0), models.OKEY).AnyTimes()


			h.Create(w, tt.args.r)
			log.Print(w.Result())

			if tt.args.status != w.Code {
				t.Error(tt.name)
			}

		})
	}
}

func TestPHandler_Read(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock.NewMockIUsecase(ctl)

	type fields struct {
		personaUsecase persona.IUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected models.PersonaRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/person/0", nil),
				expected: models.PersonaRequest{Name: "name"},
				status:   http.StatusCreated,
			}},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().Read(tt.args.expected.ID).Return(&models.PersonaResponse{}, models.OKEY).AnyTimes())

			h.Read(w, tt.args.r)

			if tt.args.status != w.Code {
				log.Print(w.Result())
			}

		})
	}
}

func TestPHandler_ReadAll(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock.NewMockIUsecase(ctl)

	type fields struct {
		personaUsecase persona.IUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected models.PersonaRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("GET", "/persons", nil),
				expected: models.PersonaRequest{Name: "name"},
				status:   http.StatusCreated,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().ReadAll().Return([]*models.PersonaResponse{}, models.OKEY).AnyTimes())

			h.ReadAll(w, tt.args.r)

			if tt.args.status != w.Code {
				log.Print(w.Result())
			}

		})
	}
}

func TestPHandler_Update(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock.NewMockIUsecase(ctl)

	type fields struct {
		personaUsecase persona.IUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected models.PersonaRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("PATCH", "/person/0", nil),
				expected: models.PersonaRequest{ID: 0},
				status:   http.StatusCreated,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().Update(tt.args.expected.ID, &tt.args.expected).Return(models.OKEY).AnyTimes())

			h.Update(w, tt.args.r)

			if tt.args.status != w.Code {
				log.Print(w.Result())
			}

		})
	}
}

func TestPHandler_Delete(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := mock.NewMockIUsecase(ctl)

	type fields struct {
		personaUsecase persona.IUsecase
	}
	type args struct {
		r        *http.Request
		result   http.Response
		status   int
		expected models.PersonaRequest
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "simple read",
			fields: fields{personaUsecase: mockUsecase},
			args: args{
				r:        httptest.NewRequest("DELETE", "/person/0", nil),
				expected: models.PersonaRequest{ID: 0},
				status:   http.StatusCreated,
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}

			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().Delete(tt.args.expected.ID).Return(models.OKEY).AnyTimes())

			h.Delete(w, tt.args.r)

			if tt.args.status != w.Code {
				t.Error()
			}

		})
	}
}
