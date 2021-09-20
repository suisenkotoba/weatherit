package events

import (
	"time"
	coordinate "weatherit/usecases/coordinates"
	"weatherit/usecases/events"

	"gorm.io/gorm"
)

type Event struct {
	ID              int
	UserID          int
	StartAt         time.Time
	EndAt           time.Time
	Title           string `gorm:"size:200"`
	Description     string
	Address         string
	GeoLat          float64
	GeoLong         float64
	EventChecklists []EventChecklist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:Restrict;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type EventChecklist struct {
	ID        int
	EventID   int
	Name      string
	IsChecked bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Event) ToDomain() events.Domain {
	return events.Domain{
		ID:             rec.ID,
		UserID:         rec.UserID,
		StartAt:        rec.StartAt,
		EndAt:          rec.EndAt,
		Title:          rec.Title,
		Description:    rec.Description,
		Address:        rec.Address,
		GeoLoc:         coordinate.CreateCoordinate([2]float64{rec.GeoLat, rec.GeoLong}),
		EventChecklist: CheckListDomain(rec.EventChecklists),
	}
}

func (rec *EventChecklist) ToDomain() events.Checklist {
	return events.Checklist{
		ID:        rec.ID,
		Name:      rec.Name,
		IsChecked: rec.IsChecked,
	}
}

func CheckListDomain(c []EventChecklist) (res []events.Checklist) {
	for i := 0; i < len(c); i++ {
		res = append(res, c[i].ToDomain())
	}
	return
}

func fromDomain(event events.Domain) *Event {
	checklists := []EventChecklist{}
	for i := 0; i < len(event.EventChecklist); i++ {
		c := EventChecklist{
			ID:        event.EventChecklist[i].ID,
			Name:      event.EventChecklist[i].Name,
			IsChecked: event.EventChecklist[i].IsChecked,
			EventID: event.ID,}
		checklists = append(checklists, c)
	}
	return &Event{
		ID:              event.ID,
		UserID:          event.UserID,
		StartAt:         event.StartAt,
		EndAt:           event.EndAt,
		Title:           event.Title,
		Description:     event.Description,
		Address:         event.Address,
		GeoLat:          event.GeoLoc.Lat,
		GeoLong:         event.GeoLoc.Long,
		EventChecklists: checklists,
	}
}
