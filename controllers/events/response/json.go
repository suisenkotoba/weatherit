package response

import (
	"weatherit/usecases/events"
)

type Checklist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsChecked bool   `json:"is_checked"`
}

type Event struct {
	ID             int         `json:"id"`
	EventDate      string      `json:"event_date"`
	StartAt        string      `json:"start_at"`
	EndAt          string      `json:"end_at"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Address        string      `json:"address"`
	GeoLoc         [2]float64  `json:"geo_loc"`
	EventChecklist []Checklist `json:"event_checklists"`
}

func FromDomain(event events.Domain) Event {
	checklist := []Checklist{}
	for i := 0; i < len(event.EventChecklist); i++ {
		checklist = append(checklist, Checklist{
			ID:        event.EventChecklist[i].ID,
			Name:      event.EventChecklist[i].Name,
			IsChecked: event.EventChecklist[i].IsChecked,
		})
	}
	return Event{
		ID:             event.ID,
		EventDate:      event.StartAt.Format("2006-01-02"),
		StartAt:        event.StartAt.Format("15:00"),
		EndAt:          event.EndAt.Format("15:00"),
		Title:          event.Title,
		Description:    event.Description,
		Address:        event.Address,
		GeoLoc:         event.GeoLoc.ToPoint(),
		EventChecklist: checklist,
	}
}
