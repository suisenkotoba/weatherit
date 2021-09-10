package events

import (
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	EventDate   time.Date
	StartAt     time.Time
	EndAt       time.Time
	Title       string
	Description string
	Address     string
	GeoLoc      Coordinate
}

type Checklist struct {
	ID 			int
	Name 		string
	IsChecked 	bool
	CheckedAt 	time.Time
}

type UseCase interface {
	GetAllUserEvents(ctx context.Context, userId int) ([]Domain, error)
	GetAllUserEventsByDateRange(ctx context.Context, userId int, from time.Date, to time.Date) ([]Domain, error)
	ScheduleEvent(ctx context.Context, event *Domain) (int, error)
	CancelEvent(ctx context.Context, eventId int) error
	UpdateEvent(ctx context.Context, event *Domain) error
	GetEventChecklist(ctx context.Context, eventID int) ([]Checklist, error)
	CreateEventCheklist(ctx context.Context, checklists []*Checklist, eventId int) (int, error)
	UpdateChecklist(ctx context.Context, checklists []*Checklist) (int, error)
	RemoveChecklist(ctx context.Context, checklistIDs []int) error
}

type Repository interface {
	Find(ctx context.Context, userId int) ([]Domain, error)
	FindByDate(ctx context.Context, userId int, from time.Date, to time.Date) ([]Domain, error)
	Store(ctx context.Context, newEvent *Domain)(int, error)
	Delete(ctx context.Context, eventId int)(int, error)
	Update(ctx context.Context, event *Domain)(int, error)
	FetchChecklist(ctx context.Context, eventId int)([]Checklist, error)
	StoreChecklist(ctx context.Context, checklist []*Checklist, eventId int) (int, error)
	UpdateChecklist(ctx context.Context, checklist []*Checklist) (int, error)
	DeleteChecklist(ctx context.Context, checklistIDs []int) error

}


type Coordinate struct {
	Lat  float
	Long float
}

func (c Coordinate) CreateCoordinate (point []int) Coordinate, error {
	if len(point) < 2 {
		return Coordinate{}, errors.New("Invalid point array")
	} else {
		return Coordinate{
			Lat: point[0],
			Long: point[1]
		}, nil
	}
}
