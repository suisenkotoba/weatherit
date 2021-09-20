package drivers

import (
	userDB "weatherit/drivers/databases/users"
	userDomain "weatherit/usecases/users"

	eventDB "weatherit/drivers/databases/events"
	eventDomain "weatherit/usecases/events"

	weatherForecasterSource "weatherit/drivers/thirdparties/weather"
	weatherForecasterDomain "weatherit/usecases/weatherforecast"

	alterplanDB "weatherit/drivers/databases/alterplans"
	alterplanDomain "weatherit/usecases/alterplan"

	activityDB "weatherit/drivers/databases/activities"
	activityDomain "weatherit/usecases/activities"

	interestDB "weatherit/drivers/databases/interests"
	interestDomain "weatherit/usecases/interests"

	userInterestDB "weatherit/drivers/databases/user_interests"
	userInterestDomain "weatherit/usecases/user_interests"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewUserRepository(conn)
}

func NewEventRepository(conn *gorm.DB) eventDomain.Repository {
	return eventDB.NewEventRepository(conn)
}

func NewWeatherForecaster(appKey string) weatherForecasterDomain.Repository {
	return weatherForecasterSource.NewWeatherForecaster(appKey)
}

func NewActivityRepository(conn *gorm.DB) activityDomain.Repository {
	return activityDB.NewActivityRepository(conn)
}

func NewAlterPlanRepository(conn *gorm.DB) alterplanDomain.Repository {
	return alterplanDB.NewAlterPlanRepository(conn)
}

func NewInterestRepository(conn *gorm.DB) interestDomain.Repository {
	return interestDB.NewInterestRepository(conn)
}

func NewUserInterestRepository(conn *gorm.DB) userInterestDomain.Repository {
	return userInterestDB.NewUserInterestRepository(conn)
}
