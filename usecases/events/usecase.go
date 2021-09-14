package events

import (
	"context"
	"time"
)

type eventUseCase struct {
	eventRepository Repository
	contextTimeout  time.Duration
}

func NewEventUseCase(timeout time.Duration, er Repository) UseCase {
	return &eventUseCase{
		eventRepository: er,
		contextTimeout:  timeout,
	}
}

func (eu *eventUseCase) GetAllUserEvents(ctx context.Context, userId int, start, end, month string) ([]Domain, error) {
	if month != "" {
		month += "-01"
		start_date, err := time.Parse("2006-01-02", month)
		if err != nil {
			return []Domain{}, err
		}
		end_date := start_date.AddDate(0, 1, -1)
		return eu.GetAllUserEventsByDateRange(ctx, userId, start_date, end_date)
	} else if start != "" && end != "" {
		start_date, err := time.Parse("2006-01-02", start)
		if err != nil {
			return []Domain{}, err
		}
		end_date, err := time.Parse("2006-01-02", end)
		if err != nil {
			return []Domain{}, err
		}
		return eu.GetAllUserEventsByDateRange(ctx, userId, start_date, end_date)
	} else {
		data, err := eu.eventRepository.Find(ctx, userId)
		if err != nil {
			return []Domain{}, err
		}
		return data, nil
	}
}

func (eu *eventUseCase) GetAllUserEventsByDateRange(ctx context.Context, userId int, from time.Time, to time.Time) ([]Domain, error) {
	data, err := eu.eventRepository.FindByDate(ctx, userId, from, to)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (eu *eventUseCase) ScheduleEvent(ctx context.Context, event *Domain) (int, error) {
	eventId, err := eu.eventRepository.Store(ctx, event)
	if err != nil {
		return 0, err
	}
	return eventId, nil
}

func (eu *eventUseCase) CancelEvent(ctx context.Context, eventId int) error {
	return nil
}

func (eu *eventUseCase) UpdateEvent(ctx context.Context, event *Domain) error {
	return nil

}

func (eu *eventUseCase) GetEventChecklist(ctx context.Context, eventID int) ([]Checklist, error) {
	return []Checklist{}, nil
}

func (eu *eventUseCase) CreateEventCheklist(ctx context.Context, checklists []*Checklist, eventId int) (int, error) {
	return 0, nil
}

func (eu *eventUseCase) UpdateChecklist(ctx context.Context, checklists []*Checklist) (int, error) {
	return 0, nil

}

func (eu *eventUseCase) RemoveChecklist(ctx context.Context, checklistIDs []int) error {
	return nil
}
