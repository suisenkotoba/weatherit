package alterplans

import (
	"context"
	"weatherit/usecases/alterplan"

	"gorm.io/gorm"
)

type mysqlAlterPlanRepository struct {
	Conn *gorm.DB
}

func NewAlterPlanRepository(conn *gorm.DB) alterplan.Repository {
	return &mysqlAlterPlanRepository{
		Conn: conn,
	}
}

func (ar *mysqlAlterPlanRepository) FindByEventID(ctx context.Context, eventId int) alterplan.Domain {
	rec := AlterPlan{}
	_ = ar.Conn.Find(&rec, "event_id = ?", eventId)
	return rec.ToDomain()

}

func (ar *mysqlAlterPlanRepository) Store(ctx context.Context, plan *alterplan.Domain) (int, error){
	rec := fromDomain(*plan)
	result := ar.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}
	return rec.ID, nil
}

func (ar *mysqlAlterPlanRepository) Update(ctx context.Context, plan *alterplan.Domain) error {
	rec := fromDomain(*plan)
	result := ar.Conn.Updates(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

