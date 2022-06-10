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
)

type Movies struct {
	m storage.Movie
}

func NewMovie(mv storage.Movie) Movies {
	return Movies{
		m: mv,
	}
}

//GetHollyWood gets hollywood movies
func (mv Movies) GetHollyWood(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)
	var ms []bson.M
	ms, err = mv.m.GetHollyWood()
	if err != nil {
		json.NewEncoder(w).Encode("No hollywood movies available")
		return
	}
	json.NewEncoder(w).Encode(ms)
}

func (mv Movies) GetLoveAndObsession(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)

	var lo []bson.M

	lo, err = mv.m.GetLoveAndObsession()
	if err != nil {
		json.NewEncoder(w).Encode("Love And Obsession movies are not available")
		return
	}
	json.NewEncoder(w).Encode(lo)
}

func (mv Movies) GetDrugLord(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)

	var res []bson.M
	res, err = mv.m.GetDrugLord()
	if err != nil {
		json.NewEncoder(w).Encode("No movie is available")
		return
	}
	json.NewEncoder(w).Encode("Movies are loading..")
	json.NewEncoder(w).Encode(res)
}

func (mv Movies) GetHollyWoodDrama(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)

	var res []bson.M
	res, err = mv.m.GetHollyWoodDrama()
	if err != nil {
		json.NewEncoder(w).Encode("No movie is available")
		return
	}
	json.NewEncoder(w).Encode("Movies are loading..")
	json.NewEncoder(w).Encode(res)

}

func (mv Movies) GetMadeInAfrica(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)

	var res []bson.M
	res, err = mv.m.GetMadeInAfrica()
	if err != nil {
		json.NewEncoder(w).Encode("No movie is available")
		return
	}
	json.NewEncoder(w).Encode("Movies are loading..")
	json.NewEncoder(w).Encode(res)

}

func (m Movies) GetMovies(w http.ResponseWriter, r *http.Request) {
	var err error
	var clm *types.Claims
	clm, err = secure.VeriyToken(w, r)
	if err != nil {
		json.NewEncoder(w).Encode("Token has expired")
	}
	g := fmt.Sprintf("Welcome Mr %s", clm.Email)
	json.NewEncoder(w).Encode(g)

	var ms []bson.M
	ms, err = m.m.GetMovies()
	if err != nil {
		json.NewEncoder(w).Encode("Collection of movies are not available")
	}
	json.NewEncoder(w).Encode("Available collection of movies")
	json.NewEncoder(w).Encode(ms)
}
