package user_interests

import (
	"context"
	"time"
)

type userInterestUseCase struct {
	userInterestRepository Repository
	contextTimeout         time.Duration
}

func NewUserInterestUseCase(timeout time.Duration, er Repository) UseCase {
	return &userInterestUseCase{
		userInterestRepository: er,
		contextTimeout:         timeout,
	}
}

func (iu *userInterestUseCase) SetUserInterest(ctx context.Context, userId int, interestIDs []int) error {
	userInterests := []Domain{}
	for i := 0; i<len(interestIDs); i++{
		userInterests = append(userInterests, 
			Domain{
				UserID: userId,
				InterestID: interestIDs[i],
			})
	}
	err := iu.userInterestRepository.Store(ctx, userInterests)
	return err
}
