package response

import (
	"time"
	"weatherit/usecases/users"
)

type Users struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	DOB     time.Time  `json:"dob"`
	Address string     `json:"address"`
	GeoLoc  [2]float64 `json:"geo_loc"`
	Gender  string     `json:"gender"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Name:    domain.Name,
		Email:  domain.Email,
		DOB:     domain.DOB,
		Address: domain.Address,
		GeoLoc:  domain.GeoLoc.ToPoint(),
		Gender:  domain.Gender,
	}
}

