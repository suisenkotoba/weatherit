package users

import (
	"net/http"
	"strings"
	controller "weatherit/controllers"
	"weatherit/controllers/interests/response"
	errorMessage "weatherit/errors"
	"weatherit/usecases/interests"

	echo "github.com/labstack/echo/v4"
)

type InterestController struct {
	interestUseCase interests.UseCase
}

func NewInterestController(ic interests.UseCase) *InterestController {
	return &InterestController{
		interestUseCase: ic,
	}
}

func (ctrl *InterestController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	dataDomain, err := ctrl.interestUseCase.GetAvailableInterests(ctx, page, limit)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	data := []response.Interest{}
	for i := 0; i < len(dataDomain); i++ {
		data = append(data, response.FromDomain(dataDomain[i]))
	}
	return controller.NewSuccessResponse(c, data)
}