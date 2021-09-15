package user_interests

import (
	"time"
	userInterests "weatherit/usecases/user_interests"

	"gorm.io/gorm"
)

type UserInterest struct {
	ID         int
	UserID     int
	InterestID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (i *UserInterest) ToDomain() userInterests.Domain {
	return userInterests.Domain{
		UserID: i.UserID,
		InterestID: i.InterestID,
	}
}

func fromDomain(ui userInterests.Domain) *UserInterest{
	return &UserInterest{
		UserID: ui.UserID,
		InterestID: ui.InterestID,
	}
}
