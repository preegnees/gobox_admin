package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (j *jsonWebToken) GenerateJWT(username string, audience string) (aToken string, rToken string, e error) {

	if username == "" || !roles.contains(audience) {
		return "", "", ErrUsernameOrRolesIsEmpty
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
	accessToken, err := at.SignedString([]byte(j.secret))
	if err != nil {
		return "", "", ErrCreateAccessToken
	}

	accessClaims.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour))
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken, err := rt.SignedString([]byte(j.secret))
	if err != nil {
		return "", "", ErrCreateRefreshToken
	}
	return accessToken, refreshToken, nil
}

func (j *jsonWebToken) CheckJwt(token string) (bool) {

	if token == "" {
		return false
	}

	_, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return false
	} else {
		return true
	}
}

func (j *jsonWebToken) GetValuesFromJWT(token string) (myClaims *MyCustomClaims, e error) {

	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
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
