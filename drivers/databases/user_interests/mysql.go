package user_interests

import (
	"context"
	userInterests "weatherit/usecases/user_interests"

	"gorm.io/gorm"
)

type mysqlUserInterestRepository struct {
	Conn *gorm.DB
}

func NewUserInterestRepository(conn *gorm.DB) userInterests.Repository {
	return &mysqlUserInterestRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserInterestRepository) Store(ctx context.Context, userInterests []userInterests.Domain) error {
	ui := []UserInterest{}
	for i:= 0; i< len(userInterests) ; i++{
		ui = append(ui, *fromDomain(userInterests[i]))
	}
	error := ur.Conn.Create(&ui).Error
	return error
}

func (ur *mysqlUserInterestRepository) FindUserInterest(ctx context.Context, userId int) []userInterests.Domain{
	ui := []UserInterest{}
	res := []userInterests.Domain{}
	err := ur.Conn.Find(&ui, "user_id = ?", userId).Error
	if err != nil {
		return res
	}
	for i:= 0; i< len(ui) ; i++{
		res = append(res, ui[i].ToDomain())
	}
	return res
}
