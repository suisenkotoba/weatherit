package routes

import (
	"encoding/json"
	"io/ioutil"
	"weatherit/controllers/events"
	"weatherit/controllers/interests"
	"weatherit/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	EventController events.EventController
	InterestController interests.InterestController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	v1 := e.Group("v1")

	v1Auth := v1.Group("/auth")
	v1Auth.POST("/register/", cl.UserController.Store)
	v1Auth.POST("/login/", cl.UserController.CreateToken)

	v1User := v1.Group("/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	v1User.GET("/me/", cl.UserController.GetProfile)
	v1User.PUT("/me/", cl.UserController.UpdateProfile)
	v1User.PUT("/loc/", cl.UserController.UpdateLocation)
	v1User.PUT("/interest/", cl.UserController.SetInterest)

	v1Interest := v1.Group("/interests", middleware.JWTWithConfig(cl.JWTMiddleware))
	v1Interest.GET("/", cl.InterestController.GetAll)

	v1Event := v1.Group("/event", middleware.JWTWithConfig(cl.JWTMiddleware))
	v1Event.POST("/", cl.EventController.Store)
	v1Event.PUT("/", cl.EventController.UpdateEvent)
	v1Event.GET("/", cl.EventController.GetEvents)
	v1Event.DELETE("/:id/", cl.EventController.DeleteEvent)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
}
