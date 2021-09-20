package users_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"
	"weatherit/app/middleware"
	errorMessage "weatherit/errors"
	"weatherit/helpers/encrypt"
	userMock "weatherit/mocks/users"
	coordinate "weatherit/usecases/coordinates"
	"weatherit/usecases/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreateToken(t *testing.T) {
	jwtAuth := middleware.ConfigJWT{
		SecretJWT:       "asdfghjkl",
		ExpiresDuration: 3600,
	}
	password, _ := encrypt.Hash("password")
	domainSample := users.Domain{
		ID:       1,
		Name:     "name",
		Email:    "email",
		Password: password,
		DOB:      time.Now(),
		Address:  "address",
		GeoLoc:   coordinate.Coordinate{Lat: -1.08, Long: 2.32},
		Gender:   "gender",
	}
	t.Run("test case 1, valid test", func(t *testing.T) {
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		token, err := userUC.CreateToken(ctx, "email", password)

		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
	t.Run("test case 2, valid test, wrong password", func(t *testing.T) {
		errorWrongPassword := errors.New(errorMessage.WrongPassword)
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		_, err := userUC.CreateToken(ctx, "email", "pasword")

		assert.Equal(t, err, errorWrongPassword)
	})
	t.Run("test case 3, valid test, password empty", func(t *testing.T) {
		errorPasswordEmpty := errors.New(errorMessage.UsernamePasswordEmpty)
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		_, err := userUC.CreateToken(ctx, "email", "")

		assert.Equal(t, err, errorPasswordEmpty)
	})
	t.Run("test case 4, error no user found", func(t *testing.T) {
		errorRepo := errors.New("Error Repo, user not found")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errorRepo)
		_, err := userUC.CreateToken(ctx, "email", "password")

		assert.NotNil(t, err)
	})
}
func TestStore(t *testing.T)          {}
func TestUpdate(t *testing.T)         {}
func TestGetByID(t *testing.T)        {}
func TestUpdateLocation(t *testing.T) {}
