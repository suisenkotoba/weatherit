package activities

import (
	"context"
	"time"
)

type activityUseCase struct {
	activityRepository Repository
	contextTimeout     time.Duration
}

func NewActivityUseCase(timeout time.Duration, ar Repository) UseCase {
	return &activityUseCase{
		activityRepository: ar,
		contextTimeout:     timeout,
	}
}

func (au *activityUseCase) GetActivitiesByInterest(ctx context.Context, interestIds []int, isOut bool) []Domain {
	return au.activityRepository.FindActivitiesByInterest(ctx, interestIds, isOut)
}

func (au *activityUseCase) GetActivitiesInOut(ctx context.Context, isOut bool) []Domain {
	return au.activityRepository.FindActivitiesInOut(ctx, isOut)
}
