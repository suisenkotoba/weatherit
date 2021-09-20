package events

import (
	"context"
	"time"

	"weatherit/usecases/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	Conn *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Repository {
	return &mysqlEventRepository{
		Conn: conn,
	}
}

func (er *mysqlEventRepository) Find(ctx context.Context, userId int) ([]events.Domain, error) {
	rec := []Event{}

	query := er.Conn

	err := query.Preload("EventChecklists").Find(&rec, "user_id = ?", userId).Error
	if err != nil {
		return []events.Domain{}, err
	}

	eventDomain := []events.Domain{}
	for _, value := range rec {
		eventDomain = append(eventDomain, value.ToDomain())
	}

	return eventDomain, nil
}

func (er *mysqlEventRepository) FindByDate(ctx context.Context, userId int, from time.Time, to time.Time) ([]events.Domain, error) {
	rec := []Event{}

	query := er.Conn

	err := query.Preload("EventChecklists").Find(&rec,
		"user_id = ? AND start_at BETWEEN ? AND ?",
		userId, from, to).Error
	if err != nil {
		return []events.Domain{}, err
	}

	eventDomain := []events.Domain{}
	for _, value := range rec {
		eventDomain = append(eventDomain, value.ToDomain())
	}

	return eventDomain, nil
}

func (er *mysqlEventRepository) FindAllByDate(ctx context.Context, from time.Time, to time.Time) ([]events.Domain, error) {
	rec := []Event{}

	query := er.Conn

	err := query.Find(&rec, "start_at BETWEEN ? AND ?", from, to).Error
	if err != nil {
		return []events.Domain{}, err
	}

	eventDomain := []events.Domain{}
	for _, value := range rec {
		eventDomain = append(eventDomain, value.ToDomain())
	}

	return eventDomain, nil
}

func (er *mysqlEventRepository) Store(ctx context.Context, newEvent *events.Domain) (int, error) {
	rec := fromDomain(*newEvent)

	result := er.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}

	return rec.ID, nil
}

func (er *mysqlEventRepository) Delete(ctx context.Context, eventId, userId int) error {
	result := er.Conn.Where("id = ? AND user_id = ?", eventId, userId).Delete(&Event{})
	return result.Error
}

func (er *mysqlEventRepository) Update(ctx context.Context, event *events.Domain) (int, error) {
	rec := fromDomain(*event)

	err := er.Conn.Updates(rec).Error

	if rec != nil {
		return 0, err
	}

	return rec.ID, nil

}
