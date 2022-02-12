package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Movie struct {
	Connec
}

func NewMovie(c Connec) Movie {
	return Movie{
		Connec: c,
	}
}

func (m Movie) GetHollyWood() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("hollywood")
	filter := bson.D{{}}
	var res []bson.M
	var num *mongo.Cursor

	num, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = num.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}

func (m Movie) GetLoveAndObsession() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("loveobsession")
	filter := bson.D{{}}
	var res []bson.M
	var count *mongo.Cursor
	count, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	if err = count.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (m Movie) GetDrugLord() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("druglord")
	var co *mongo.Cursor
	filter := bson.D{{}}

	co, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var res []bson.M
	if err = co.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (m Movie) GetHollyWoodDrama() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("hollywooddrama")
	var co *mongo.Cursor
	filter := bson.D{{}}

	co, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var res []bson.M
	if err = co.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}

func (m Movie) GetMadeInAfrica() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("madeInAfrica")
	var co *mongo.Cursor
	filter := bson.D{{}}

	co, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var res []bson.M
	if err = co.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}

func (m Movie) GetMovies() ([]bson.M, error) {
	var err error
	db := m.Client.Database("mydb").Collection("allmovies")
	filter := bson.D{{}}
	var curs *mongo.Cursor
	curs, err = db.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var res []bson.M
	if err = curs.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil

}
