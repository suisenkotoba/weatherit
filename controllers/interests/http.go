package interests

import (
	"net/http"
	controller "weatherit/controllers"
	"weatherit/controllers/interests/response"
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

	// page := c.QueryParam("page")
	// limit := c.QueryParam("limit")

	// pageSize, offset := helpers.OffsetLimit(page, limit)
	dataDomain, err := ctrl.interestUseCase.GetAvailableInterests(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	data := []response.Interest{}
	for i := 0; i < len(dataDomain); i++ {
		data = append(data, response.FromDomain(dataDomain[i]))
	}
	return controller.NewSuccessResponse(c, data)
}
