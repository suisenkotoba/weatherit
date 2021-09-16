package events

import (
	"context"
	"fmt"
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

type mysqlEventCheklistRepository struct {
	Conn *gorm.DB
}

func NewEventChecklistRepository(conn *gorm.DB) events.ChecklistRepository {
	return &mysqlEventCheklistRepository{
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

func (er *mysqlEventRepository) FindAllByDate(ctx context.Context, from time.Time, to time.Time) ([]events.Domain, error){
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

func (er *mysqlEventRepository) Delete(ctx context.Context, eventId, userId int) (int, error) {
	result := er.Conn.Where("id = ? AND user_id = ?", eventId, userId).Delete(&Event{})
	fmt.Println(result)
	return 0, result.Error
}

func (er *mysqlEventRepository) Update(ctx context.Context, event *events.Domain) (int, error) {
	rec := fromDomain(*event)

	checklistRec := rec.EventChecklists
	rec.EventChecklists = []EventChecklist{}

	tx := er.Conn.Begin()
	eventResult := tx.Save(&rec)
	if eventResult.Error != nil {
		tx.Rollback()
		return 0, eventResult.Error
	}

	for i := 0; i< len(checklistRec); i++{
		checklistResult := tx.Save(&checklistRec[i])
		if checklistResult.Error != nil {
			tx.Rollback()
			return 0, checklistResult.Error
		}
	}
	
	tx.Commit()

	return rec.ID, nil
}

func (evr *mysqlEventCheklistRepository) Fetch(ctx context.Context, eventId int) ([]events.Checklist, error) {
	return []events.Checklist{}, nil
}

func (evr *mysqlEventCheklistRepository) Store(ctx context.Context, checklist []*events.Checklist, eventId int) (int, error) {
	return 0, nil
}

func (evr *mysqlEventCheklistRepository) Update(ctx context.Context, checklist []*events.Checklist) (int, error) {
	return 0, nil
}

func (evr *mysqlEventCheklistRepository) Delete(ctx context.Context, checklistIDs []int) error {
	return nil
}
