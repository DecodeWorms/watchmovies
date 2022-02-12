package secure

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"watchmovies/types"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string) (*types.TokenDetails, error) {
	var err error
	token := &types.TokenDetails{}
	token.Aexp = time.Now().Add(time.Minute * 15)
	token.Rexp = time.Now().Add(time.Hour * 24 * 7)

	aclaims := &types.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token.Aexp.Unix(),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, aclaims)
	token.AccessToken, err = tkn.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		panic(err)
	}

	rclaims := &types.Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: token.Rexp.Unix(),
		},
	}
	rtkn := jwt.NewWithClaims(jwt.SigningMethodHS256, rclaims)
	token.RefreshToken, err = rtkn.SignedString([]byte("REFRESH_SECRET"))
	if err != nil {
		panic(err)
	}
	return token, nil

}

func SetCookie(w http.ResponseWriter, r *http.Request, n string) {

	cookie := &http.Cookie{
		Name:   "token",
		Value:  n,
		MaxAge: 180,
	}
	http.SetCookie(w, cookie)
}

func VeriyToken(w http.ResponseWriter, r *http.Request) (*types.Claims, error) {
	c, _ := r.Cookie("token")
	token := c.Value

	clm := &types.Claims{}

	tkn, _ := jwt.ParseWithClaims(token, clm, func(v *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if !tkn.Valid {
		fmt.Println("Token has been xpired")
	}
	return clm, nil
}
