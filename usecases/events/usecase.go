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
	return &categoryUsecase{
		eventRepository: er,
		contextTimeout:  timeout,
	}
}

func (eu *eventUseCase) GetAllUserEvents(ctx context.Context, userId int) ([]Domain, error) {
	data, err := eu.eventRepository.Find(userId)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (eu *eventUseCase) GetAllUserEventsByDateRange(ctx context.Context, userId int, from time.Date, to time.Date) ([]Domain, error) {
	data, err := eu.eventRepository.FindByDate(userId, from, to)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (eu *eventUseCase) ScheduleEvent(ctx context.Context, event *Domain) (int, error) {
	eventId, err := eu.eventRepository.Store(event)
	if err != nil {
		return 0, err
	}
	return eventId, nil
}

func (eu *eventUseCase) CancelEvent(ctx context.Context, eventId int) error {

}

func (eu *eventUseCase) UpdateEvent(ctx context.Context, event *Domain) error {

}

func (eu *eventUseCase) GetEventChecklist(ctx context.Context, eventID int) ([]Checklist, error) {

}

func (eu *eventUseCase) CreateEventCheklist(ctx context.Context, checklists []*Checklist, eventId int) (int, error) {
}

func (eu *eventUseCase) UpdateChecklist(ctx context.Context, checklists []*Checklist) (int, error) {

}

func (eu *eventUseCase) RemoveChecklist(ctx context.Context, checklistIDs []int) error {

}
