package jwt

import (
	"errors"
)

var ErrUsernameOrRolesIsEmpty = errors.New("Err Username Or Roles Is Empty")
var ErrCreateAccessToken = errors.New("Err Create Access Token")
var ErrCreateRefreshToken = errors.New("Err Create Refresh Token")
var ErrParseWithClaims = errors.New("Err Parse With Claims")
var ErrGetClaimsFromToken = errors.New("Err Get Claims From Token")
var ErrGetEnvSecret = errors.New("Err Get Env Secret")