package types

type Users struct {
	Email       string
	Password    string
	Gender      string
	Country     string
	Plan        int
	CreditCards CreditCards
}

type CreditCards struct {
	CardType       string
	CardNumber     int
	AccountBalance int
	ExpireTime     string
	SecurityPin    int
}
