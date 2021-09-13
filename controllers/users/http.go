package users

import (
	"net/http"
	"strings"
	controller "weatherit/controllers"
	"weatherit/controllers/users/request"
	errorMessage "weatherit/errors"
	"weatherit/usecases/users"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(uc users.UseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.userUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			return controller.NewErrorResponse(c, http.StatusBadRequest, err)
		} else {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *UserController) CreateToken(c echo.Context) error {
	ctx := c.Request().Context()

	loginReq := request.UserLogin{}
	if err := c.Bind(&loginReq); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userUseCase.CreateToken(ctx, loginReq.Email, loginReq.Password)
	if err != nil {
		if err.Error() == errorMessage.WrongPassword {
			return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"accessToken"`
	}{Token: token}

	return controller.NewSuccessResponse(c, response)
}
