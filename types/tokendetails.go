package types

import "time"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	Aexp         time.Time
	Rexp         time.Time
}

type Tkn struct{}
