package user_interests

import (
	"time"
	"weatherit/drivers/databases/interests"
	"weatherit/drivers/databases/users"
	userInterests "weatherit/usecases/user_interests"

	"gorm.io/gorm"
)

type UserInterest struct {
	ID         int
	UserID     int
	InterestID int
	User       users.User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Interest   interests.Interest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (i *UserInterest) ToDomain() userInterests.Domain {
	return userInterests.Domain{
		UserID:     i.UserID,
		InterestID: i.InterestID,
	}
}

func fromDomain(ui userInterests.Domain) *UserInterest {
	return &UserInterest{
		UserID:     ui.UserID,
		InterestID: ui.InterestID,
	}
}
