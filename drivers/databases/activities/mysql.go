package activities

import (
	"context"
	"weatherit/usecases/activities"

	"gorm.io/gorm"
)

type mysqlActivityRepository struct {
	Conn *gorm.DB
}

func NewActivityRepository(conn *gorm.DB) activities.Repository {
	return &mysqlActivityRepository{
		Conn: conn,
	}
}

func (ar *mysqlActivityRepository) FindActivitiesByInterest(ctx context.Context, interestIds []int) []activities.Domain {
	rec := []Activity{}
	err := ar.Conn.Preload("Interest").Find(&rec, "interest_id IN (?)", interestIds).Error
	if err != nil {
		return []activities.Domain{}
	}
	domains := []activities.Domain{}
	for i:=0; i<len(rec) ; i++{
		domains = append(domains, rec[i].ToDomain())
	}
	return domains
}

func (ar *mysqlActivityRepository) FindActivitiesInOut(ctx context.Context, isOut bool) []activities.Domain {
	rec := []Activity{}
	err := ar.Conn.Preload("Interest").Find(&rec, "is_outdoor = ?", isOut).Error
	if err != nil {
		return []activities.Domain{}
	}
	domains := []activities.Domain{}
	for i:=0; i<len(rec) ; i++{
		domains = append(domains, rec[i].ToDomain())
	}
	return domains
}
