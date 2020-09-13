package delivery

import (
	"RSOI/internal/models"
	"RSOI/internal/pkg/persona"
	"RSOI/internal/pkg/persona/mock"
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
		r      *http.Request
		result http.Response
		status int
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
					strings.NewReader("")),
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PHandler{
				personaUsecase: tt.fields.personaUsecase,
			}
			w := httptest.NewRecorder()

			gomock.InOrder(
				mockUsecase.EXPECT().Create(models.PersonaRequest{}).Return(uint(0), models.OKEY).AnyTimes())

			h.Create(w, tt.args.r)

			if w.Code != tt.args.status {
				log.Print(w.Result())
			}

		})
	}
}
