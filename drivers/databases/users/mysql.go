package users

import (
	"context"
	"weatherit/usecases/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) GetByID(ctx context.Context, userId int) (users.Domain, error){
	rec := User{}
	err := ur.Conn.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (ur *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := User{}
	err := ur.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.ToDomain(), nil
}

func (ur *mysqlUserRepository) Store(ctx context.Context, data *users.Domain) (int, error) {
	rec := fromDomain(*data)

	result := ur.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}

	return rec.ID, nil
}

func (ur *mysqlUserRepository) Update(ctx context.Context, data *users.Domain) error{
	rec := fromDomain(*data)
	result := ur.Conn.Updates(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
