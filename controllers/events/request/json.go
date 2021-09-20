package request

import (
	"weatherit/helpers/datetime"
	coordinate "weatherit/usecases/coordinates"
	"weatherit/usecases/events"
)

type Event struct {
	EventDate      string     `json:"event_date"`
	StartAt        string     `json:"start_at"`
	EndAt          string     `json:"end_at"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Address        string     `json:"address"`
	GeoLoc         [2]float64 `json:"geo_loc"`
	EventChecklist []string   `json:"event_checklists"`
}

type Checklist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsChecked bool   `json:"is_checked"`
}

type UpdatedEvent struct {
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

func (req *Event) ToDomain(userID int) (*events.Domain, error) {
	startAt, err := datetime.CombineDateTime(req.EventDate, req.StartAt+":00")
	if err != nil {
		return &events.Domain{}, err
	}
	endAt, err := datetime.CombineDateTime(req.EventDate, req.EndAt+":00")
	if err != nil {
		return &events.Domain{}, err
	}
	return &events.Domain{
		UserID:         userID,
		StartAt:        startAt,
		EndAt:          endAt,
		Title:          req.Title,
		Description:    req.Description,
		Address:        req.Address,
		GeoLoc:         coordinate.CreateCoordinate(req.GeoLoc),
		EventChecklist: checklistDomainFromStr(req.EventChecklist),
	}, nil
}

func (req *UpdatedEvent) ToDomain(userID int) (*events.Domain, error) {
	startAt, err := datetime.CombineDateTime(req.EventDate, req.StartAt+":00")
	if err != nil {
		return &events.Domain{}, err
	}
	endAt, err := datetime.CombineDateTime(req.EventDate, req.EndAt+":00")
	if err != nil {
		return &events.Domain{}, err
	}
	return &events.Domain{
		ID:             req.ID,
		UserID:         userID,
		StartAt:        startAt,
		EndAt:          endAt,
		Title:          req.Title,
		Description:    req.Description,
		Address:        req.Address,
		GeoLoc:         coordinate.CreateCoordinate(req.GeoLoc),
		EventChecklist: checklistDomain(req.EventChecklist),
	}, nil
}

func checklistDomain(checklist []Checklist) (res []events.Checklist) {
	for i := 0; i < len(checklist); i++ {
		res = append(res, events.Checklist{
			ID:        checklist[i].ID,
			Name:      checklist[i].Name,
			IsChecked: checklist[i].IsChecked,
		})
	}
	return
}

func checklistDomainFromStr(checklist []string) (res []events.Checklist) {
	for i := 0; i < len(checklist); i++ {
		res = append(res, events.Checklist{
			Name:      checklist[i],
			IsChecked: false,
		})
	}
	return
}
