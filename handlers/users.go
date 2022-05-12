package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"watchmovies/secure"
	"watchmovies/storage"
	"watchmovies/types"
	"watchmovies/util"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlers struct {
	us storage.UsersStorage
}

func NewUserHandlers(u storage.UsersStorage) UserHandlers {
	return UserHandlers{
		us: u,
	}

}
func (u UserHandlers) Create(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var user types.Users
	var err error
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON object")
		return

	}
	err = u.us.Create(user)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to create resources")
		return
	}
	json.NewEncoder(w).Encode("Resources created successfully")

}

func (u UserHandlers) Login(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var user types.Users
	_ = json.NewDecoder(r.Body).Decode(&user)
	var tk *types.TokenDetails
	tk, _ = secure.CreateToken(user.Email)
	secure.SetCookie(w, r, tk.AccessToken)
	var res bson.M
	res, err = u.us.Login(user)
	f := res["password"]
	str := fmt.Sprintf("%v", f)

	if err = bcrypt.CompareHashAndPassword([]byte(str), []byte(user.Password)); err != nil {
		json.NewEncoder(w).Encode("Either email or password is incorrect")
		return
	}
	json.NewEncoder(w).Encode("User Logged in successfully")
	fmt.Println("hello, world")
	fmt.Println("hello, world2")
	fmt.Println("commit one")
	fmt.Println("commit two")
	fmt.Println("commit three")

}
