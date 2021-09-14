package interests

import (
	"time"
	"weatherit/usecases/interests"

	"gorm.io/gorm"
)

type Interest struct {
	ID                 int
	Name               string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

func (i *Interest) ToDomain() interests.Domain {
	return interests.Domain{
		ID: i.ID,
		Name: i.Name,
	}
}