package jwt

import "os"

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IJWT interface {
	GenerateJWT(username string, audience string) (aToken string, rToken string, e error)
	CheckJwt(token string) (bool)
	GetValuesFromJWT(token string) (myClaims *MyCustomClaims, e error)
}

type jsonWebToken struct {
	secret string
}

func New() (IJWT, error) {
	
	secret := os.Getenv("SECRET")
	if secret == "" {
		return nil, ErrGetEnvSecret 
	}
	return &jsonWebToken{
		secret: secret,
	}, nil
}