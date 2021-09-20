package users

import (
	"net/http"
	"strings"
	"weatherit/app/middleware"
	controller "weatherit/controllers"
	"weatherit/controllers/users/request"
	"weatherit/controllers/users/response"
	errorMessage "weatherit/errors"
	userInterests "weatherit/usecases/user_interests"
	"weatherit/usecases/users"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase         users.UseCase
	userInterestUseCase userInterests.UseCase
}

func NewUserController(uc users.UseCase, ic userInterests.UseCase) *UserController {
	return &UserController{
		userUseCase:         uc,
		userInterestUseCase: ic,
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

func (ctrl *UserController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)

	profile, err := ctrl.userUseCase.GetByID(ctx, user.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(profile))
}

func (ctrl *UserController) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)

	profile := request.Users{}
	if err := c.Bind(&profile); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	userprofile := profile.ToDomain()
	userprofile.ID = user.ID
	err := ctrl.userUseCase.Update(ctx, userprofile)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "Profile updated!")
}

func (ctrl *UserController) UpdateLocation(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)

	newLoc := request.UserLocation{}
	if err := c.Bind(&newLoc); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := ctrl.userUseCase.UpdateLocation(ctx, user.ID, newLoc.GeoLoc[0], newLoc.GeoLoc[1])
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "Location updated!")
}

func (ctrl *UserController) SetInterest(c echo.Context) error {
	ctx := c.Request().Context()
	user := middleware.GetUser(c)

	interests := request.UserInterests{}
	if err := c.Bind(&interests); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	err := ctrl.userInterestUseCase.SetUserInterest(ctx, user.ID, interests.Interests)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, "Successfully add interest(s)")
}
