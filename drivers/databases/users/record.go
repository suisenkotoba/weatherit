package users

import (
	"time"
	coordinate "weatherit/usecases/coordinates"
	"weatherit/usecases/users"

	"gorm.io/gorm"
)

type User struct {
	ID        int
	Name      string `gorm:"size:200"`
	Email     string `gorm:"uniqueIndex;size:200"`
	Password  string
	DOB       time.Time
	Address   string
	GeoLat    float64
	GeoLong   float64
	Gender    string `gorm:"size:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:       rec.ID,
		Name:     rec.Name,
		Email:    rec.Email,
		Password: rec.Password,
		DOB:      rec.DOB,
		Address:  rec.Address,
		GeoLoc:   coordinate.CreateCoordinate([2]float64{rec.GeoLat, rec.GeoLong}),
		Gender:   rec.Gender,
	}
}

func fromDomain(userDomain users.Domain) *User {
	return &User{
		ID:       userDomain.ID,
		Name:     userDomain.Name,
		Email:    userDomain.Email,
		Password: userDomain.Password,
		DOB:      userDomain.DOB,
		Address:  userDomain.Address,
		GeoLat:   userDomain.GeoLoc.Lat,
		GeoLong:  userDomain.GeoLoc.Long,
		Gender:   userDomain.Gender,
	}
}
