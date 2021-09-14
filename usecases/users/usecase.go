package users

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"weatherit/app/middleware"
	errorMessage "weatherit/errors"
	"weatherit/helpers/encrypt"
)

type userUseCase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUseCase(er Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) UseCase {
	return &userUseCase{
		userRepository: er,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (uc *userUseCase) CreateToken(ctx context.Context, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", errors.New(errorMessage.UsernamePasswordEmpty)
	}

	userDomain, err := uc.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", errors.New(errorMessage.WrongPassword)
	}

	token := uc.jwtAuth.GenerateToken(userDomain.ID, userDomain.Email)
	return token, nil
}

func (uc *userUseCase) Store(ctx context.Context, data *Domain) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByEmail(ctx, data.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return 0, err
		}
	}
	if existedUser != (Domain{}) {
		return 0, errors.New(fmt.Sprintf(errorMessage.EmailAlreadyRegistered, data.Email))
	}
	data.Password, err = encrypt.Hash(data.Password)
	if err != nil {
		return 0, errors.New(errorMessage.SomethingWrong)
	}
	id, err := uc.userRepository.Store(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *userUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return uc.userRepository.GetByID(ctx, id)
}

func (uc *userUseCase) Update(ctx context.Context, data *Domain) error {
	return uc.userRepository.Update(ctx, data)
}

func (uc *userUseCase) UpdateLocation(ctx context.Context, userId int, lat float64, long float64) error {
	user, err := uc.GetByID(ctx, userId)
	if err != nil {
		return err
	}
	user.GeoLoc.Lat = lat
	user.GeoLoc.Long = long

	err = uc.Update(ctx, &user)
	return err

}
