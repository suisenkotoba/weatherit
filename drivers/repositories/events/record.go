package events

import "time"

type Event struct {
	ID          int
	UserID      int
	EventDate   time.Date
	StartAt     time.Time
	EndAt       time.Time
	Title       string
	Description string
	Address     string
}

func (rec *Event) ToDomain() events.Domain {
	return events.Domain{
		ID: rec.ID
		UserID: rec.UserID
		EventDate: rec.EventDate
		StartAt: rec.StartAt
		EndAt: rec.EndAt
		Title: rec.Title
		Description: rec.Description
		Address : rec.Address
	}
}
