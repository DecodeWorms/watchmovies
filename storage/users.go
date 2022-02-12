package storage

import (
	"context"
	"watchmovies/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UsersStorage struct {
	Connec
}

func NewStorage(s Connec) UsersStorage {
	return UsersStorage{
		Connec: s,
	}

}

// type Plans struct {
// 	HDClear     int
// 	HDmoreClear int
// 	HDmostClear int
// }

func (s UsersStorage) Create(use types.Users) error {
	var err error
	db := s.Connec.Client.Database("mydb").Collection("users")
	pass, _ := bcrypt.GenerateFromPassword([]byte(use.Password), 12)

	u := types.Users{
		Email:    use.Email,
		Password: string(pass),
		Gender:   use.Gender,
		Country:  use.Country,
		Plan:     use.Plan,
		CreditCards: types.CreditCards{
			CardType:       use.CreditCards.CardType,
			CardNumber:     use.CreditCards.CardNumber,
			AccountBalance: use.CreditCards.AccountBalance,
			ExpireTime:     use.CreditCards.ExpireTime,
			SecurityPin:    use.CreditCards.SecurityPin,
		},
	}
	_, err = db.InsertOne(context.TODO(), u)
	if err != nil {
		return nil
	}
	return err

}

func (s UsersStorage) Login(use types.Users) (bson.M, error) {
	var err error
	db := s.Client.Database("mydb").Collection("users")
	filter := bson.D{{"email", use.Email}}
	proj := bson.D{{"email", 1}, {"password", 1}}
	opts := options.FindOne().SetProjection(proj)
	var res bson.M
	if err = db.FindOne(context.TODO(), filter, opts).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil

}
