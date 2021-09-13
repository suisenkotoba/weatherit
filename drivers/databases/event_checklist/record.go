package events

import "time"

type EventChecklist struct {
	ID        int
	EvetID    int
	Name      string
	IsChecked bool
	CheckedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
