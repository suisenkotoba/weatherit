package alterplan

import "context"

type Domain struct {
	ID                int
	EventID           int
	ActivityID        int
	WeatherForecastH1 string
	WeatherForecastH6 string
	WeatherForecastD1 string
	IsTaken           bool
}

type UseCase interface {
	GetEventAlterPlan(ctx context.Context, eventId int) Domain
	MakeEventAlterPlan(ctx context.Context, plan *Domain) (int, error)
	UpdateEventAlterPlan(ctx context.Context, plan *Domain) error
}

type Repository interface {
	FindByEventID(ctx context.Context, eventId int)  Domain
	Store(ctx context.Context, plan *Domain) (int, error)
	Update(ctx context.Context, plan *Domain) error
}
