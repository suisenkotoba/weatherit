package drivers

import (
	userDB "weatherit/drivers/databases/users"
	userDomain "weatherit/usecases/users"

	eventDB "weatherit/drivers/databases/events"
	eventDomain "weatherit/usecases/events"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewUserRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewEventRepository(conn)
}
