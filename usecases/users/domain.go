package users

import (
	"context"
	"time"
	coordinate "weatherit/usecases/coordinates"
)

type Domain struct {
	ID       int
	Name     string
	Email    string
	Password string
	DOB      time.Time
	Address  string
	GeoLoc   coordinate.Coordinate
	Gender   string
}

type UseCase interface {
	CreateToken(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *Domain) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	UpdateLocation(ctx context.Context, userId int, lat float64, long float64)
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) (int, error)
	Update(ctx context.Context, data *Domain)
}
