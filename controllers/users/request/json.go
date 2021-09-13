package request

import (
	"time"
	coordinate "weatherit/usecases/coordinates"
	"weatherit/usecases/users"
)

type Users struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	DOB      string     `json:"dob"`
	Address  string     `json:"address"`
	GeoLoc   [2]float64 `json:"geo_loc"`
	Gender   string     `json:"gender"`
}

func (req *Users) ToDomain() *users.Domain {
	DOB, _ := time.Parse("2006-01-02", req.DOB)
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		DOB:      DOB,
		Address:  req.Address,
		GeoLoc:   coordinate.CreateCoordinate(req.GeoLoc),
		Gender:   req.Gender,
	}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
