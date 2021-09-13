package main

import (
	"log"
	"time"
	_middleware "weatherit/app/middleware"
	_routes "weatherit/app/routes"
	_driverFactory "weatherit/drivers"
	_dbDriver "weatherit/drivers/databases"

	_userController "weatherit/controllers/users"
	_userUsecase "weatherit/usecases/users"

	_eventController "weatherit/controllers/events"
	_eventUsecase "weatherit/usecases/events"

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

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUseCase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	eventRepo := _driverFactory.NewEventRepository(db)
	eventUsecase := _eventUsecase.NewEventUseCase(timeoutContext, eventRepo)
	eventCtrl := _eventController.NewEventController(eventUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:   configJWT.Init(),
		UserController:  *userCtrl,
		EventController: *eventCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
