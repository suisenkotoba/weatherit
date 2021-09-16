package activities

import (
	"time"

	"weatherit/drivers/databases/interests"
	"weatherit/usecases/activities"

	"gorm.io/gorm"
)

type Activity struct {
	ID                 int
	Name               string
	IsOutdoor          bool
	RecommendedWeather string
	InterestID         int
	Interest           interests.Interest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

func (rec *Activity) ToDomain() activities.Domain {
	return activities.Domain{
		ID: rec.ID,
		Name: rec.Name,
		IsOutdoor: rec.IsOutdoor,
		RecommendedWeather: rec.RecommendedWeather,
		InterestID: rec.InterestID,
	}
}
