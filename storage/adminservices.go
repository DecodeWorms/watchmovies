package storage

import (
	"context"
	"watchmovies/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type Services struct {
	Connec
}

func NewServices(c Connec) Services {
	return Services{
		Connec: c,
	}
}

func (s Services) HollyWood(m types.HollyWood) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("hollywood")
	mo := types.HollyWood{
		Greetings:      m.Greetings,
		Title:          m.Title,
		Image:          m.Image,
		Video:          m.Video,
		TimeDuration:   m.TimeDuration,
		YearReleased:   m.YearReleased,
		AgeRestriction: m.AgeRestriction,
		Likes:          m.Likes,
	}

	var num *mongo.InsertOneResult
	num, err = db.InsertOne(context.TODO(), mo)
	if err != nil {
		return nil, err
	}
	return num, nil

}

func (s Services) LoveAndObsession(love types.LoveAndObsession) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("loveobsession")
	mcol := types.LoveAndObsession{
		Greetings:      love.Greetings,
		Title:          love.Title,
		Image:          love.Image,
		Video:          love.Video,
		TimeDuration:   love.TimeDuration,
		YearReleased:   love.YearReleased,
		AgeRestriction: love.AgeRestriction,
		Likes:          love.Likes,
	}
	var cur *mongo.InsertOneResult
	cur, err = db.InsertOne(context.TODO(), mcol)
	if err != nil {
		return nil, err
	}
	return cur, nil

}

func (s Services) DrugLord(drug types.DrugLord) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("druglord")
	mcol := types.LoveAndObsession{
		Greetings:      drug.Greetings,
		Title:          drug.Title,
		Image:          drug.Image,
		Video:          drug.Video,
		TimeDuration:   drug.TimeDuration,
		YearReleased:   drug.YearReleased,
		AgeRestriction: drug.AgeRestriction,
		Likes:          drug.Likes,
	}
	var cur *mongo.InsertOneResult
	cur, err = db.InsertOne(context.TODO(), mcol)
	if err != nil {
		return nil, err
	}
	return cur, nil

}
func (s Services) HollyWoodDrama(holly types.HollyWoodDrama) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("hollywooddrama")
	mcol := types.HollyWoodDrama{
		Greetings:      holly.Greetings,
		Title:          holly.Title,
		Image:          holly.Image,
		Video:          holly.Video,
		TimeDuration:   holly.TimeDuration,
		YearReleased:   holly.YearReleased,
		AgeRestriction: holly.AgeRestriction,
		Likes:          holly.Likes,
	}
	var cur *mongo.InsertOneResult
	cur, err = db.InsertOne(context.TODO(), mcol)
	if err != nil {
		return nil, err
	}
	return cur, nil

}

func (s Services) MadeInAfrica(africa types.MadeInAfrica) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("madeInAfrica")
	mcol := types.MadeInAfrica{
		Greetings:      africa.Greetings,
		Title:          africa.Title,
		Image:          africa.Image,
		Video:          africa.Video,
		TimeDuration:   africa.TimeDuration,
		YearReleased:   africa.YearReleased,
		AgeRestriction: africa.AgeRestriction,
		Likes:          africa.Likes,
	}
	var cur *mongo.InsertOneResult
	cur, err = db.InsertOne(context.TODO(), mcol)
	if err != nil {
		return nil, err
	}
	return cur, nil

}

func (s Services) MovieCat(cat types.Categories) (*mongo.InsertOneResult, error) {
	var err error
	db := s.Client.Database("mydb").Collection("categories")

	categ := types.Categories{
		MoviesSelection: cat.MoviesSelection,
		HollyWood: types.HollyWood{
			Title:          cat.HollyWood.Title,
			Image:          cat.HollyWood.Image,
			Video:          cat.HollyWood.Video,
			TimeDuration:   cat.HollyWood.TimeDuration,
			YearReleased:   cat.HollyWood.YearReleased,
			AgeRestriction: cat.HollyWood.AgeRestriction,
			Likes:          cat.HollyWood.Likes,
		},
		LoveAndObsession: types.LoveAndObsession{
			Title:          cat.LoveAndObsession.Title,
			Image:          cat.LoveAndObsession.Image,
			Video:          cat.LoveAndObsession.Video,
			TimeDuration:   cat.LoveAndObsession.TimeDuration,
			YearReleased:   cat.LoveAndObsession.YearReleased,
			AgeRestriction: cat.LoveAndObsession.AgeRestriction,
			Likes:          cat.LoveAndObsession.Likes,
		},
		DrugLord: types.DrugLord{
			Title:          cat.DrugLord.Title,
			Image:          cat.DrugLord.Image,
			Video:          cat.DrugLord.Video,
			TimeDuration:   cat.DrugLord.TimeDuration,
			YearReleased:   cat.DrugLord.YearReleased,
			AgeRestriction: cat.DrugLord.AgeRestriction,
			Likes:          cat.DrugLord.Likes,
		},
		HollyWoodDrama: types.HollyWoodDrama{
			Title:          cat.HollyWoodDrama.Title,
			Image:          cat.HollyWoodDrama.Image,
			Video:          cat.HollyWoodDrama.Video,
			TimeDuration:   cat.HollyWoodDrama.TimeDuration,
			YearReleased:   cat.HollyWoodDrama.YearReleased,
			AgeRestriction: cat.HollyWoodDrama.AgeRestriction,
			Likes:          cat.HollyWoodDrama.Likes,
		},
		MadeInAfrica: types.MadeInAfrica{
			Title:          cat.MadeInAfrica.Title,
			Image:          cat.MadeInAfrica.Image,
			Video:          cat.MadeInAfrica.Video,
			TimeDuration:   cat.MadeInAfrica.TimeDuration,
			YearReleased:   cat.MadeInAfrica.YearReleased,
			AgeRestriction: cat.MadeInAfrica.AgeRestriction,
			Likes:          cat.MadeInAfrica.Likes,
		},
	}

	var num *mongo.InsertOneResult
	num, err = db.InsertOne(context.TODO(), categ)
	if err != nil {
		return nil, err
	}
	return num, nil

}
