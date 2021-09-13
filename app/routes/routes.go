package routes

import (
	"encoding/json"
	"io/ioutil"
	"weatherit/controllers/events"
	"weatherit/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	EventController events.EventController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	v1 := e.Group("v1")

	v1Auth := v1.Group("/auth")
	v1Auth.POST("/register/", cl.UserController.Store)
	v1Auth.POST("/login/", cl.UserController.CreateToken)

	v1Event := v1.Group("/event", middleware.JWTWithConfig(cl.JWTMiddleware))
	v1Event.POST("/", cl.EventController.Store)
	v1Event.GET("/", cl.EventController.GetEvents)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
}
