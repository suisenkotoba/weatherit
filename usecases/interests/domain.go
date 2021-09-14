package interests

import "context"

type Domain struct {
	ID                 int
	Name               string
}

type UseCase interface {
	GetAvailableInterests(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
}
