package events

import (
	"context"

	"gorm.io/gorm"
)

type eventRepository struct {
	conn *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Repository {
	return &eventRepository{
		conn: conn,
	}
}

func (er *eventRepository) Find(ctx context.Context, userId int) ([]events.Domain, error) {
	rec := []Event{}

	query := cr.conn.Where("user_id = ?", userId)

	err := query.Find(&rec).Error
	if err != nil {
		return []event.Domain{}, err
	}

	eventDomain := []event.Domain{}
	for _, value := range rec {
		eventDomain = append(eventDomain, value.ToDomain())
	}

	return eventDomain, nil
}
