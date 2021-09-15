package user_interests

import "context"

type Domain struct {
	UserID     int
	InterestID int
}

type UseCase interface {
	SetUserInterest(ctx context.Context, userId int, interestIDs []int) error
}

type Repository interface {
	Store(ctx context.Context, userInterests []Domain) error
}
