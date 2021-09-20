package alterplans

import (
	"time"
	"weatherit/drivers/databases/activities"
	"weatherit/drivers/databases/events"
	"weatherit/usecases/alterplan"
	null "gopkg.in/guregu/null.v4"
)

type AlterPlan struct {
	ID                int
	EventID           int
	ActivityID        null.Int
	WeatherForecastH1 string
	WeatherForecastH6 string
	WeatherForecastD1 string
	IsTaken           bool
	Event             events.Event        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Activity          activities.Activity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (rec *AlterPlan) ToDomain() alterplan.Domain {
	return alterplan.Domain{
		ID:                rec.ID,
		EventID:           rec.EventID,
		ActivityID:        int(rec.ActivityID.ValueOrZero()),
		WeatherForecastH1: rec.WeatherForecastH1,
		WeatherForecastH6: rec.WeatherForecastH6,
		WeatherForecastD1: rec.WeatherForecastD1,
	}
}

func fromDomain(domain alterplan.Domain) *AlterPlan {
	activityID := null.Int{}
	if domain.ActivityID != 0 {
		activityID = null.IntFrom(int64(domain.ActivityID))
	}
	return &AlterPlan{
		ID:                domain.ID,
		EventID:           domain.EventID,
		ActivityID:        activityID,
		WeatherForecastH1: domain.WeatherForecastH1,
		WeatherForecastH6: domain.WeatherForecastH6,
		WeatherForecastD1: domain.WeatherForecastD1,
	}
}
