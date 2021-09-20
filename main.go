package main

import (
	"log"
	"os"
	"time"
	_middleware "weatherit/app/middleware"
	_routes "weatherit/app/routes"
	_driverFactory "weatherit/drivers"
	_dbDriver "weatherit/drivers/databases"
	"weatherit/jobs"

	_userController "weatherit/controllers/users"
	_users "weatherit/usecases/users"

	_eventController "weatherit/controllers/events"
	_activities "weatherit/usecases/activities"
	_alterplans "weatherit/usecases/alterplan"
	_events "weatherit/usecases/events"

	_interestController "weatherit/controllers/interests"
	_interests "weatherit/usecases/interests"

	_userInterests "weatherit/usecases/user_interests"

	"weatherit/usecases"

	"github.com/jasonlvhit/gocron"
	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	opt := "api"
	if len(os.Args) > 1 {
		opt = os.Args[1]
	}

	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _users.NewUserUseCase(userRepo, &configJWT, timeoutContext)
	userInterestRepo := _driverFactory.NewUserInterestRepository(db)
	userInterestUsecase := _userInterests.NewUserInterestUseCase(timeoutContext, userInterestRepo)

	eventRepo := _driverFactory.NewEventRepository(db)
	weatherForecaster := _driverFactory.NewWeatherForecaster(viper.GetString(`openweather.key`))
	eventUsecase := _events.NewEventUseCase(timeoutContext, eventRepo, weatherForecaster)

	activityRepo := _driverFactory.NewActivityRepository(db)
	activityUsecase := _activities.NewActivityUseCase(timeoutContext, activityRepo)

	alterplanRepo := _driverFactory.NewAlterPlanRepository(db)
	alterplanUsecase := _alterplans.NewAlterPlanUseCase(timeoutContext, alterplanRepo)

	interestRepo := _driverFactory.NewInterestRepository(db)
	interestUsecase := _interests.NewInterestUseCase(timeoutContext, interestRepo)

	if opt == "api" {
		e := echo.New()

		userCtrl := _userController.NewUserController(userUsecase, userInterestUsecase)
		eventCtrl := _eventController.NewEventController(eventUsecase)
		interestCtrl := _interestController.NewInterestController(interestUsecase)

		routesInit := _routes.ControllerList{
			JWTMiddleware:      configJWT.Init(),
			UserController:     *userCtrl,
			EventController:    *eventCtrl,
			InterestController: *interestCtrl,
		}

		routesInit.RouteRegister(e)
		log.Fatal(e.Start(viper.GetString("server.address")))
	} else if opt == "scheduler" {
		ucList := usecases.UsecaseList{
			Event:        eventUsecase,
			User:         userUsecase,
			Interest:     interestUsecase,
			UserInterest: userInterestUsecase,
			AlterPlan:    alterplanUsecase,
			Activity:     activityUsecase,
		}
		scheduler := gocron.NewScheduler()
		scheduler.Every(1).Minute().Do(jobs.Forecast, ucList, "H1")
		scheduler.Every(1).Minute().Do(jobs.Forecast, ucList, "H6")
		scheduler.Every(1).Minute().Do(jobs.Forecast, ucList, "D1")
		<-scheduler.Start()
	}
}
