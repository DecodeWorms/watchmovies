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

type MainCard struct {
	Id      int    `json:"id"`
	Country string `json:"country"`
}
type RealCard struct {
	Id        int     `json:"id"`
	MaxAmount float32 `json:"max_amount"`
}

type OutsideCard struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type VerifyCard struct {
	Id           int    `json:"id"`
	SerialNumber string `json:"serial_number"`
	Address      string `json:"address"`
	Sibling      string `json:"sibling"`
}
