package events

import (
	"context"
	"time"
	coordinate "weatherit/usecases/coordinates"
)

type Domain struct {
	ID             int
	UserID         int
	StartAt        time.Time
	EndAt          time.Time
	Title          string
	Description    string
	Address        string
	GeoLoc         coordinate.Coordinate
	EventChecklist []Checklist
}

type Checklist struct {
	ID        int
	Name      string
	IsChecked bool
}

type UseCase interface {
	GetAllUserEvents(ctx context.Context, userId int, from, to, month string) ([]Domain, error)
	GetAllUserEventsByDateRange(ctx context.Context, userId int, from time.Time, to time.Time) ([]Domain, error)
	ScheduleEvent(ctx context.Context, event *Domain) (int, error)
	CancelEvent(ctx context.Context, eventId, userId int) error
	UpdateEvent(ctx context.Context, event *Domain) error
	GetEventChecklist(ctx context.Context, eventID int) ([]Checklist, error)
	CreateEventCheklist(ctx context.Context, checklists []*Checklist, eventId int) (int, error)
	UpdateChecklist(ctx context.Context, checklists []*Checklist) (int, error)
	RemoveChecklist(ctx context.Context, checklistIDs []int) error
}

type Repository interface {
	Find(ctx context.Context, userId int) ([]Domain, error)
	FindByDate(ctx context.Context, userId int, from time.Time, to time.Time) ([]Domain, error)
	Store(ctx context.Context, newEvent *Domain) (int, error)
	Delete(ctx context.Context, eventId, userId int) (int, error)
	Update(ctx context.Context, event *Domain) (int, error)
}

type ChecklistRepository interface {
	Fetch(ctx context.Context, eventId int) ([]Checklist, error)
	Store(ctx context.Context, checklist []*Checklist, eventId int) (int, error)
	Update(ctx context.Context, checklist []*Checklist) (int, error)
	Delete(ctx context.Context, checklistIDs []int) error
}
