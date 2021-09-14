package interests

import (
	"context"
	"time"
)

type interestUseCase struct {
	interestRepository Repository
	contextTimeout     time.Duration
}

func NewInterestUseCase(timeout time.Duration, er Repository) UseCase {
	return &interestUseCase{
		interestRepository: er,
		contextTimeout:     timeout,
	}
}

func (iu *interestUseCase) GetAvailableInterests(ctx context.Context) ([]Domain, error) {
	data, err := iu.interestRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}
