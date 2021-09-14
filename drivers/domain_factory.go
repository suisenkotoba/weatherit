package drivers

import (
	userDB "weatherit/drivers/databases/users"
	userDomain "weatherit/usecases/users"

	eventDB "weatherit/drivers/databases/events"
	eventDomain "weatherit/usecases/events"

	interestDB "weatherit/drivers/databases/interests"
	interestDomain "weatherit/usecases/interests"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewUserRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewEventRepository(conn)
}

func NewInterestRepository(conn *gorm.DB) interestDomain.Repository {
	return interestDB.NewInterestRepository(conn)
}
