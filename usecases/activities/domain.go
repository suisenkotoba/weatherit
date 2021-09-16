package activities

import "context"

type Domain struct {
	ID                 int
	Name               string
	IsOutdoor          bool
	RecommendedWeather string
	InterestID         int
}

type UseCase interface {
	GetActivitiesByInterest(ctx context.Context, interestIds []int) []Domain
	GetActivitiesInOut(ctx context.Context, isOut bool) []Domain
}

type Repository interface {
	FindActivitiesByInterest(ctx context.Context, interestIds []int) []Domain
	FindActivitiesInOut(ctx context.Context, isOut bool) []Domain
}
