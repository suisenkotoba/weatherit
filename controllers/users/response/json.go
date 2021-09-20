package response

import (
	"weatherit/usecases/users"
)

type Users struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	DOB     string     `json:"dob"`
	Address string     `json:"address"`
	GeoLoc  [2]float64 `json:"geo_loc"`
	Gender  string     `json:"gender"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Name:    domain.Name,
		Email:   domain.Email,
		DOB:     domain.DOB.Format("2006-01-02"),
		Address: domain.Address,
		GeoLoc:  domain.GeoLoc.ToPoint(),
		Gender:  domain.Gender,
	}
}
