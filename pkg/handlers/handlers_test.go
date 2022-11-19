package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "jwt/pkg/models"
	serviceMock "jwt/pkg/services/mock"
	jwtMock "jwt/pkg/utils/jwt/mock"
	validator_ "jwt/pkg/utils/validator"

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
	e.Validator = &validator_.CustomValidator{
		Validator: validator.New(),
	}

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-in", strings.NewReader(d.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req_ := httptest.NewRequest(http.MethodPost, "/auth/sign-in", strings.NewReader(d.json))
		req_.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec_ := httptest.NewRecorder()
		c_ := e.NewContext(req_, rec_)

		signInTest := models.SignIn{}
		if err := c_.Bind(&signInTest); err != nil {
			panic(err)
		}

		service.EXPECT().SignIn(models.SignIn{
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

func TestSignUp(t *testing.T) {

	data := []struct {
		json      string
		status    int
		returnErr error
	}{
		{
			json:      `{"name":"Jon Snow","email":"jon@labstack.com"}`,
			status:    http.StatusBadRequest,
			returnErr: nil,
		},
		{
			json:      `{"username":"username","password":"password", "role":"Admin", "email":"123@gmial.com", "role": "User", "email_code": 1}`,
			status:    http.StatusOK,
			returnErr: nil,
		},
		{
			json:      `{"username":"username","password":"password", "role":"Admin", "email":"123@gmial.com", "role": "User", "email_code": 1}`,
			status:    http.StatusInternalServerError,
			returnErr: errors.New("test"),
		},
	}

	e := echo.New()
	e.Validator = &validator_.CustomValidator{
		Validator: validator.New(),
	}

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", strings.NewReader(d.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req_ := httptest.NewRequest(http.MethodPost, "/auth/sign-up", strings.NewReader(d.json))
		req_.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec_ := httptest.NewRecorder()
		c_ := e.NewContext(req_, rec_)

		signUpTest := models.SignUp{}
		if err := c_.Bind(&signUpTest); err != nil {
			panic(err)
		}

		service.EXPECT().SignUp(models.SignUp{
			Username:  signUpTest.Username,
			Password:  signUpTest.Password,
			Role:      signUpTest.Role,
			Email:     signUpTest.Email,
			EmailCode: signUpTest.EmailCode,
		}).Return(d.returnErr).AnyTimes()

		if err := handlers.AuthSignUp(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}

func TestSignOut(t *testing.T) {

	data := []struct {
		name      string
		value     string
		status    int
		returnErr error
	}{
		{
			name:      "refresh_token",
			value:     "",
			status:    http.StatusBadRequest,
			returnErr: nil,
		},
		{
			name:      "refresh_token",
			value:     "2.2.2",
			status:    http.StatusOK,
			returnErr: nil,
		},
		{
			name:      "refresh_token",
			value:     "3.3.3",
			status:    http.StatusInternalServerError,
			returnErr: errors.New("test"),
		},
	}

	e := echo.New()

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		cookie := new(http.Cookie)
		cookie.Name = d.name
		cookie.Value = d.value
		cookie.HttpOnly = true
		req.AddCookie(cookie)

		c := e.NewContext(req, rec)

		service.EXPECT().SignOut(d.value).Return(d.returnErr).AnyTimes()

		if err := handlers.AuthSignOut(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}

func TestAuthRefresh(t *testing.T) {

	data := []struct {
		name      string
		value     string
		status    int
		returnErr error
	}{
		{
			name:      "refresh_token",
			value:     "",
			status:    http.StatusBadRequest,
			returnErr: nil,
		},
		{
			name:      "refresh_token",
			value:     "2.2.2",
			status:    http.StatusOK,
			returnErr: nil,
		},
		{
			name:      "refresh_token",
			value:     "3.3.3",
			status:    http.StatusInternalServerError,
			returnErr: errors.New("test"),
		},
	}

	e := echo.New()

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/auth/sign-out", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		cookie := new(http.Cookie)
		cookie.Name = d.name
		cookie.Value = d.value
		cookie.HttpOnly = true
		req.AddCookie(cookie)

		c := e.NewContext(req, rec)

		service.EXPECT().SignOut(d.value).Return(d.returnErr).AnyTimes()

		if err := handlers.AuthSignOut(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}

func TestApiSaveAppData(t *testing.T) {

	data := []struct {
		json      string
		status    int
		returnErr error
	}{
		{
			json:      `{"username":"username","tokens":[{"token":"token1", "action": 1}]}`,
			status:    http.StatusOK,
			returnErr: nil,
		},
		{
			json:      ``,
			status:    http.StatusBadRequest,
			returnErr: nil,
		},
		{
			json:      `{"username":"username","tokens":[{"token":"token1", "action": 1}]}`,
			status:    http.StatusInternalServerError,
			returnErr: errors.New("..."),
		},
	}

	e := echo.New()
	e.Validator = &validator_.CustomValidator{
		Validator: validator.New(),
	}

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/api/save", strings.NewReader(d.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req_ := httptest.NewRequest(http.MethodPost, "/api/save", strings.NewReader(d.json))
		req_.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec_ := httptest.NewRecorder()
		c_ := e.NewContext(req_, rec_)

		saveData := models.AppData{}
		if err := c_.Bind(&saveData); err != nil {
			panic(err)
		}

		service.EXPECT().SaveAppData(saveData.Username, saveData.Tokens).Return(d.returnErr).AnyTimes()

		if err := handlers.ApiSaveAppData(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}

func TestApiGetAppData(t *testing.T) {

	data := []struct {
		json         string
		status       int
		returnTokens []models.Tokens
		returnErr    error
	}{
		{
			json:   `{"username":"username"}`,
			status: http.StatusOK,
			returnTokens: []models.Tokens{
				{
					Token:  "token1",
					Action: 0,
				},
				{
					Token:  "token2",
					Action: 0,
				},
			},
			returnErr: nil,
		},
		{
			json:         ``,
			status:       http.StatusBadRequest,
			returnTokens: []models.Tokens{},
			returnErr:    nil,
		},
		{
			json:         `{"username":"username"}`,
			status:       http.StatusInternalServerError,
			returnTokens: []models.Tokens{},
			returnErr:    errors.New("..."),
		},
	}

	e := echo.New()
	e.Validator = &validator_.CustomValidator{
		Validator: validator.New(),
	}

	for _, d := range data {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		service := serviceMock.NewMockIService(mockCtrl)
		jwt_ := jwtMock.NewMockIJWT(mockCtrl)

		handlers := New(service, jwt_)

		req := httptest.NewRequest(http.MethodPost, "/api/get", strings.NewReader(d.json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req_ := httptest.NewRequest(http.MethodPost, "/api/get", strings.NewReader(d.json))
		req_.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec_ := httptest.NewRecorder()
		c_ := e.NewContext(req_, rec_)

		getData := models.AppData{}
		if err := c_.Bind(&getData); err != nil {
			panic(err)
		}

		service.EXPECT().GetAppData(getData.Username).Return(d.returnTokens, d.returnErr).AnyTimes()

		if err := handlers.ApiGetAppData(c); err != nil {
			panic(err)
		}
		if rec.Result().StatusCode != d.status {
			panic("rec.Result().StatusCode != status")
		}
	}
}
