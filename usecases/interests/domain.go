package interests

import "context"

type Domain struct {
	ID                 int
	Name               string
}

type UseCase interface {
	GetAvailableInterests(ctx context.Context, limit, offset int) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, limit, offset int) ([]Domain, error)
}
