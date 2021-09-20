package interests_test

import (
	"context"
	"errors"
	"os"
	"testing"
	interestMock "weatherit/mocks/interests"
	interests "weatherit/usecases/interests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetAvailableInterests(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		var interestRepo interestMock.Repository
		interestsUC := interests.NewInterestUseCase(2, &interestRepo)
		interestDomain := []interests.Domain{
			{
				ID:   1,
				Name: "name",
			},
		}
		interestRepo.On("Find", mock.Anything).Return(interestDomain, nil).Once()
		ctx := context.Background()
		result, err := interestsUC.GetAvailableInterests(ctx)

		assert.Nil(t, err)
		assert.Equal(t, interestDomain, result)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		repoError := errors.New("Error Repo")
		var interestRepo interestMock.Repository
		interestsUC := interests.NewInterestUseCase(2, &interestRepo)
		interestRepo.On("Find", mock.Anything).Return([]interests.Domain{}, repoError).Once()
		ctx := context.Background()
		_, err := interestsUC.GetAvailableInterests(ctx)

		assert.Equal(t, repoError, err)
	})
}
