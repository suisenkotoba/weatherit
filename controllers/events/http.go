package events

import (
	"net/http"
	"weatherit/app/middleware"
	controller "weatherit/controllers"
	"weatherit/controllers/events/request"
	"weatherit/controllers/events/response"

	// errorMessage "weatherit/errors"
	"weatherit/usecases/events"

	echo "github.com/labstack/echo/v4"
)

type EventController struct {
	eventUseCase events.UseCase
}

func NewEventController(ec events.UseCase) *EventController {
	return &EventController{
		eventUseCase: ec,
	}
}

func (ctrl *EventController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)
	req := request.Event{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	newEvent, err := req.ToDomain(user.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	_, err2 := ctrl.eventUseCase.ScheduleEvent(ctx, newEvent)
	if err2 != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully shceduled event")
}

func (ctrl *EventController) GetEvents(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)

	start := c.QueryParam("start")
	end := c.QueryParam("end")
	month := c.QueryParam("month")

	dataDomain, err := ctrl.eventUseCase.GetAllUserEvents(ctx, user.ID, start, end, month)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	data := []response.Event{}
	for i := 0; i < len(dataDomain); i++ {
		data = append(data, response.FromDomain(dataDomain[i]))
	}
	return controller.NewSuccessResponse(c, data)
}
