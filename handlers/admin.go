package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"watchmovies/storage"
	"watchmovies/types"
	"watchmovies/util"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	serv storage.Services
}

func NewServices(s storage.Services) Services {
	return Services{
		serv: s,
	}
}

//Admin activities

func (s Services) HollyWood(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var mov types.HollyWood
	if err = json.NewDecoder(r.Body).Decode(&mov); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var n *mongo.InsertOneResult
	n, err = s.serv.HollyWood(mov)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to create resources(s)")
	}
	json.NewEncoder(w).Encode("Resource created successfully")
	json.NewEncoder(w).Encode(n)
}

func (s Services) LoveAndObsession(w http.ResponseWriter, r *http.Request) {
	var err error
	var love types.LoveAndObsession
	if err = json.NewDecoder(r.Body).Decode(&love); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var num *mongo.InsertOneResult
	num, err = s.serv.LoveAndObsession(love)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	json.NewEncoder(w).Encode(num)
}

func (s Services) DrugLord(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var drug types.DrugLord
	if err = json.NewDecoder(r.Body).Decode(&drug); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var num *mongo.InsertOneResult
	num, err = s.serv.DrugLord(drug)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	json.NewEncoder(w).Encode("Resource created successfully")
	json.NewEncoder(w).Encode(num)

}

func (s Services) Hollywooddrama(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var hollyw types.HollyWoodDrama
	if err = json.NewDecoder(r.Body).Decode(&hollyw); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var num *mongo.InsertOneResult
	num, err = s.serv.HollyWoodDrama(hollyw)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	json.NewEncoder(w).Encode("Resource created successfully")
	json.NewEncoder(w).Encode(num)

}

func (s Services) MadeInAfrica(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var africa types.MadeInAfrica
	if err = json.NewDecoder(r.Body).Decode(&africa); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	var num *mongo.InsertOneResult
	num, err = s.serv.MadeInAfrica(africa)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON")
		return
	}
	json.NewEncoder(w).Encode(num)

}

func (s Services) MovieCat(w http.ResponseWriter, r *http.Request) {
	util.SetHeader(w)
	var err error
	var cat types.Categories
	if err = json.NewDecoder(r.Body).Decode(&cat); err != nil {
		json.NewEncoder(w).Encode("Unable to Unmarshal JSON...")
		return
	}
	var num *mongo.InsertOneResult
	num, err = s.serv.MovieCat(cat)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to create resource(s)")
		return
	}
	json.NewEncoder(w).Encode("Resources are created successfully")
	json.NewEncoder(w).Encode(num)
	fmt.Println("hello, world")

}
