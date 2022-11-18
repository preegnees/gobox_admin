package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ms "jwt/pkg/models"
	se "jwt/pkg/services/mock"
	jt "jwt/pkg/utils/jwt/mock"

	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestSignIn(t *testing.T) {

	data := []struct {
		json                    string
		status                  int
		returnOkSignIn          bool
		returnErrSignIn         error
		returnATokenGenerateJWT string
		returnRTokenGenerateJWT string
		returnErrGenerateJWT    error
	}{
		{
			json:                    `{"name":"Jon Snow","email":"jon@labstack.com"}`,
			status:                  http.StatusBadRequest,
			returnOkSignIn:          false,
			returnErrSignIn:         nil,
			returnATokenGenerateJWT: "access test1",
			returnRTokenGenerateJWT: "refresh test1",
			returnErrGenerateJWT:    nil,
		},
		{
			json:                    `{"username":"username","password":"password", "role":"Admin"}`,
			status:                  http.StatusOK,
			returnOkSignIn:          true,
			returnErrSignIn:         nil,
			returnATokenGenerateJWT: "access test1",
			returnRTokenGenerateJWT: "refresh test1",
			returnErrGenerateJWT:    nil,
		},
		{
			json:                    `{"username":"username","password":"password", "role":"Admin"}`,
			status:                  http.StatusUnauthorized,
			returnOkSignIn:          false,
			returnErrSignIn:         nil,
			returnATokenGenerateJWT: "",
			returnRTokenGenerateJWT: "",
			returnErrGenerateJWT:    nil,
		},
		{
			json:                    `{"username":"username","password":"password", "role":"Admin"}`,
			status:                  http.StatusInternalServerError,
			returnOkSignIn:          true,
			returnErrSignIn:         errors.New("test"),
			returnATokenGenerateJWT: "",
			returnRTokenGenerateJWT: "",
			returnErrGenerateJWT:    nil,
		},
	}

	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := se.NewMockIService(mockCtrl)
		jwt_ := jt.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-in", strings.NewReader(d.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req_ := httptest.NewRequest(http.MethodPost, "/auth/sign-in", strings.NewReader(d.json))
		req_.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec_ := httptest.NewRecorder()
		c_ := e.NewContext(req_, rec_)

		signInTest := ms.SignIn{}
		if err := c_.Bind(&signInTest); err != nil {
			panic(err)
		}

		service.EXPECT().SignIn(ms.SignIn{
			Username: signInTest.Username,
			Password: signInTest.Password,
			Role:     signInTest.Role,
		}).Return(d.returnOkSignIn, d.returnErrSignIn).AnyTimes()

		jwt_.EXPECT().GenerateJWT(
			signInTest.Username, signInTest.Role,
		).Return(d.returnATokenGenerateJWT, d.returnRTokenGenerateJWT, d.returnErrGenerateJWT).AnyTimes()

		if err := handlers.AuthSignIn(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}
