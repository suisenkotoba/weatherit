package user_interests_test

import (
	"context"
	"errors"
	"os"
	"testing"
	userInterestMock "weatherit/mocks/user_interests"
	userInterests "weatherit/usecases/user_interests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSetUserInterest(t *testing.T){
	t.Run("test case 1, valid test", func(t *testing.T) {
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestRepo.On("Store", mock.Anything, mock.AnythingOfType("[]user_interests.Domain")).Return(nil).Once()
		interestsIDs := []int{1,2,3,4}
		err := userInterestUC.SetUserInterest(ctx, 1, interestsIDs)
		assert.Nil(t, err)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		repoError := errors.New("Error Repo")
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestRepo.On("Store", mock.Anything, mock.AnythingOfType("[]user_interests.Domain")).Return(repoError).Once()
		interestsIDs := []int{1,2,3,4}
		err := userInterestUC.SetUserInterest(ctx, 1, interestsIDs)
		assert.Equal(t, err, repoError)
	})
}

func TestGetUserInterest(t *testing.T){
	t.Run("test case 1, valid test", func(t *testing.T) {
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestDomain := []userInterests.Domain{
			{
				UserID: 1,
				InterestID: 1,
			},
		}
		userInterestRepo.On("FindUserInterest", mock.Anything, mock.AnythingOfType("int")).Return(userInterestDomain)
		res := userInterestUC.GetUserInterest(ctx, 1)
		assert.NotEmpty(t, res)
	})
	t.Run("test case 2, error raised", func(t *testing.T) {
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestRepo.On("FindUserInterest", mock.Anything, mock.AnythingOfType("int")).Return([]userInterests.Domain{})
		res := userInterestUC.GetUserInterest(ctx, 1)
		assert.Equal(t, len(res), 0)
	})
}

func TestGetUserInterestIDs(t *testing.T){
	t.Run("test case 1, valid test", func(t *testing.T) {
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestDomain := []userInterests.Domain{
			{
				UserID: 1,
				InterestID: 1,
			},
		}
		supposedIDs := []int{userInterestDomain[0].InterestID}
		userInterestRepo.On("FindUserInterest", mock.Anything, mock.AnythingOfType("int")).Return(userInterestDomain)
		res := userInterestUC.GetUserInterestIDs(ctx, 1)
		assert.Equal(t, supposedIDs, res)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		var userInterestRepo userInterestMock.Repository
		ctx := context.Background()
		userInterestUC := userInterests.NewUserInterestUseCase(2, &userInterestRepo)
		userInterestRepo.On("FindUserInterest", mock.Anything, mock.AnythingOfType("int")).Return([]userInterests.Domain{})
		res := userInterestUC.GetUserInterestIDs(ctx, 1)
		assert.Empty(t, res)
	})
}