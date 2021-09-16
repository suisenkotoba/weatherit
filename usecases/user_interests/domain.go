package user_interests

import "context"

type Domain struct {
	UserID     int
	InterestID int
}

type UseCase interface {
	SetUserInterest(ctx context.Context, userId int, interestIDs []int) error
	GetUserInterest(ctx context.Context, userId int) []Domain
	GetUserInterestIDs(ctx context.Context, userId int) []int
}

type Repository interface {
	Store(ctx context.Context, userInterests []Domain) error
	FindUserInterest(ctx context.Context, userId int) []Domain
}
