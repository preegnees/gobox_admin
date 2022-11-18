package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	Username string `json:"username"`
}

type CustomClaims struct {
	MyCustomClaims
	jwt.RegisteredClaims
}

const Admin string = "Admin"
const User string = "User"

type Roles []string
func (r *Roles) contains(s string) bool {
	for _, v := range *r {
		if s == v {
			return true
		}
	}
	return false
}
var roles Roles = []string{Admin, User}