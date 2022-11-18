package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const Admin string = "Admin"
const User string = "User"

var ErrUsernameOrSecretIsEmpty = errors.New("Err Username Or Secret Is Empty")
var ErrCreateAccessToken = errors.New("Err Create Access Token")
var ErrCreateRefreshToken = errors.New("Err Create Refresh Token")
var ErrParseWithClaims = errors.New("Err Parse With Claims")
var ErrGetClaimsFromToken = errors.New("Err Get Claims From Token")

type MyCustomClaims struct {
	Username string `json:"username"`
}

type CustomClaims struct {
	MyCustomClaims
	jwt.RegisteredClaims
}

func GenerateJWT(secret string, username string, audience string) (string, string, error) {

	if secret == "" || username == "" {
		return "", "", ErrUsernameOrSecretIsEmpty
	}

	accessClaims := CustomClaims{
		MyCustomClaims{
			Username: username,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(10 * time.Minute)),
			Issuer:    "gobox",
			Subject:   "auth",
			Audience:  jwt.ClaimStrings{audience},
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", "", ErrCreateAccessToken
	}

	accessClaims.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour))
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken, err := rt.SignedString([]byte(secret))
	if err != nil {
		return "", "", ErrCreateRefreshToken
	}
	return accessToken, refreshToken, nil
}

func CheckJwt(secret string, token string) bool {

	if token == "" {
		return false
	}

	_, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetValuesFromJWT(secret string, token string) (*MyCustomClaims, error) {

	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrParseWithClaims
	}

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return &MyCustomClaims{
			Username: claims.Username,
		}, nil
	} else {
		return nil, ErrGetClaimsFromToken
	}
}
