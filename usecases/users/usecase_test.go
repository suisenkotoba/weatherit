package users_test

import (
	"context"
	"errors"
	"fmt"
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
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		token, err := userUC.CreateToken(ctx, "email", "password")

		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
	t.Run("test case 2, valid test, wrong password", func(t *testing.T) {
		errorWrongPassword := errors.New(errorMessage.WrongPassword)
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		_, err := userUC.CreateToken(ctx, "email", "pasword")

		assert.Equal(t, err, errorWrongPassword)
	})
	t.Run("test case 3, valid test, password empty", func(t *testing.T) {
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		_, err := userUC.CreateToken(ctx, "email", "")

		assert.Equal(t, err.Error(), errorMessage.UsernamePasswordEmpty)
	})
	t.Run("test case 4, error no user found", func(t *testing.T) {
		errorRepo := errors.New("Error Repo, user not found")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errorRepo)
		_, err := userUC.CreateToken(ctx, "email", "password")

		assert.NotNil(t, err)
	})
}
func TestStore(t *testing.T) {
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
		errorNoUser := errors.New("Error Repo, user not found")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errorNoUser)
		userRepo.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(1, nil)
		newUserDomain := users.Domain{
			Name:     "name",
			Email:    "email",
			Password: "hello",
			DOB:      time.Now(),
			Address:  "address",
			GeoLoc:   coordinate.Coordinate{Lat: -1.08, Long: 2.32},
			Gender:   "gender",
		}
		userID, err := userUC.Store(ctx, &newUserDomain)

		assert.Nil(t, err)
		assert.NotEmpty(t, userID)
	})

	t.Run("test case 2, email already registered", func(t *testing.T) {
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domainSample, nil)
		newUserDomain := users.Domain{
			Name:     "name",
			Email:    "email",
			Password: "hello",
			DOB:      time.Now(),
			Address:  "address",
			GeoLoc:   coordinate.Coordinate{Lat: -1.08, Long: 2.32},
			Gender:   "gender",
		}
		_, err := userUC.Store(ctx, &newUserDomain)

		expectedErrMessage := fmt.Sprintf(errorMessage.EmailAlreadyRegistered, domainSample.Email)
		assert.Equal(t, err.Error(), expectedErrMessage)
	})
	t.Run("test case 3, error raised", func(t *testing.T) {
		errorNoUser := errors.New("Error Repo, user not found")
		errorStoreRepo := errors.New("Error store repo")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errorNoUser)
		userRepo.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(0, errorStoreRepo)
		newUserDomain := users.Domain{
			Name:     "name",
			Email:    "email",
			Password: "hello",
			DOB:      time.Now(),
			Address:  "address",
			GeoLoc:   coordinate.Coordinate{Lat: -1.08, Long: 2.32},
			Gender:   "gender",
		}
		_, err := userUC.Store(ctx, &newUserDomain)
		assert.NotNil(t, err)
	})

	t.Run("test case 4, failed to check user exists", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errorRepo)
		newUserDomain := users.Domain{
			Name:     "name",
			Email:    "email",
			Password: "hello",
			DOB:      time.Now(),
			Address:  "address",
			GeoLoc:   coordinate.Coordinate{Lat: -1.08, Long: 2.32},
			Gender:   "gender",
		}
		_, err := userUC.Store(ctx, &newUserDomain)

		assert.NotNil(t, err)
	})
}
func TestUpdate(t *testing.T) {
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
	t.Run("test case 1, valid case", func(t *testing.T) {
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil)
		err := userUC.Update(ctx, &domainSample)
		assert.Nil(t, err)
	})
	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(errorRepo)
		err := userUC.Update(ctx, &domainSample)
		assert.NotNil(t, err)
	})
}
func TestGetByID(t *testing.T) {
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
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domainSample, nil)

		_, err := userUC.GetByID(ctx, 1)

		assert.Nil(t, err)

	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errorRepo)

		_, err := userUC.GetByID(ctx, 1)

		assert.NotNil(t, err)

	})
}
func TestUpdateLocation(t *testing.T) {
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
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domainSample, nil)
		userRepo.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil)

		err := userUC.UpdateLocation(ctx, 1, 2.1232, -0.232)

		assert.Nil(t, err)

	})

	t.Run("test case 2, error user not found", func(t *testing.T) {
		errorRepo := errors.New("Error Repo, user not found")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errorRepo)
		userRepo.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(errorRepo)

		err := userUC.UpdateLocation(ctx, 1, 2.1232, -0.232)

		assert.NotNil(t, err)

	})
	t.Run("test case 3, error update", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var userRepo userMock.Repository
		ctx := context.Background()
		userUC := users.NewUserUseCase(&userRepo, &jwtAuth, 2)
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domainSample, nil)
		userRepo.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(errorRepo)

		err := userUC.UpdateLocation(ctx, 1, 2.1232, -0.232)

		assert.NotNil(t, err)

	})

}
