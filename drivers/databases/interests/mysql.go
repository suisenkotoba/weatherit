package interests

import (
	"context"
	"weatherit/usecases/interests"

	"gorm.io/gorm"
)

type mysqlInterestRepository struct {
	Conn *gorm.DB
}

func NewInterestRepository(conn *gorm.DB) interests.Repository {
	return &mysqlInterestRepository{
		Conn: conn,
	}
}

func (ir *mysqlInterestRepository) Find(ctx context.Context, limit, offset int) ([]interests.Domain, error) {
	rec := []Interest{}

	query := ir.Conn

	err := query.Find(&rec).Offset(offset).Limit(limit).Error
	if err != nil {
		return []interests.Domain{}, err
	}

	interestDomain := []interests.Domain{}
	for _, value := range rec {
		interestDomain = append(interestDomain, value.ToDomain())
	}

	return interestDomain, nil
}
