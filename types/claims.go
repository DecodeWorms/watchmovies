package types

import "github.com/dgrijalva/jwt-go"

//claim doc has comment on claim object
type Claims struct {
	Email string
	jwt.StandardClaims
}
