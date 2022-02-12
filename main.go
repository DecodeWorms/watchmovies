package main

import (
	"log"
	"net/http"
	"watchmovies/handlers"
	"watchmovies/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *storage.Connec
var c *mongo.Client
var sto storage.UsersStorage
var ser storage.Services
var mv storage.Movie

func init() {
	_ = godotenv.Load()
	client = storage.NewConnect(c)
	sto = storage.NewStorage(*client)
	ser = storage.NewServices(*client)
	mv = storage.NewMovie(*client)

}

func main() {
	router := mux.NewRouter()

	userHand := handlers.NewUserHandlers(sto)
	router.HandleFunc("/users/create", userHand.Create).Methods("POST")
	router.HandleFunc("/users/login", userHand.Login).Methods(http.MethodGet)

	//Match admin incoming request and dispach an handler for it
	services := handlers.NewServices(ser)
	router.HandleFunc("/services/hollywood", services.HollyWood).Methods(http.MethodPost)
	router.HandleFunc("/services/loveobsesee", services.LoveAndObsession).Methods(http.MethodPost)
	router.HandleFunc("/services/druglord", services.DrugLord).Methods(http.MethodPost)
	router.HandleFunc("/services/hollywooddrama", services.Hollywooddrama).Methods(http.MethodPost)
	router.HandleFunc("/services/madeinafrica", services.MadeInAfrica).Methods(http.MethodPost)
	router.HandleFunc("/services/catmovies", services.MovieCat).Methods(http.MethodPost)

	//Match client incoming request and dispatch an handler
	movies := handlers.NewMovie(mv)
	router.HandleFunc("/users/movies/hollywood", movies.GetHollyWood).Methods(http.MethodGet)
	router.HandleFunc("/users/movies/druglord", movies.GetDrugLord).Methods(http.MethodGet)
	router.HandleFunc("/users/movies/hollywooddrama", movies.GetHollyWoodDrama).Methods(http.MethodGet)
	router.HandleFunc("/users/movies/madeinafrica", movies.GetMadeInAfrica).Methods(http.MethodGet)
	router.HandleFunc("/users/movies/loveobsess", movies.GetLoveAndObsession).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8001", router))

}
