package alterplan

import (
	"context"
	"time"
)

type alterPlanUseCase struct {
	alterPlanRepository Repository
	contextTimeout      time.Duration
}

func NewAlterPlanUseCase(timeout time.Duration, ar Repository) UseCase {
	return &alterPlanUseCase{
		alterPlanRepository: ar,
		contextTimeout:      timeout,
	}
}

func (ac *alterPlanUseCase) GetEventAlterPlan(ctx context.Context, eventId int) Domain {
	plan := ac.alterPlanRepository.FindByEventID(ctx, eventId)
	return plan
}

func (ac *alterPlanUseCase) MakeEventAlterPlan(ctx context.Context, plan *Domain) (int, error) {
	return ac.alterPlanRepository.Store(ctx, plan)
}

func (ac *alterPlanUseCase) UpdateEventAlterPlan(ctx context.Context, plan *Domain) error {
	return ac.alterPlanRepository.Update(ctx, plan)
}
