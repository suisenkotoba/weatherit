package activities_test

import (
	"context"
	"os"
	"testing"
	activityMock "weatherit/mocks/activities"
	activities "weatherit/usecases/activities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetActivitiesByInterest(t *testing.T){
	t.Run("test case 1, valid test", func(t *testing.T) {
		var activityRepo activityMock.Repository
		ctx := context.Background()
		activityUC := activities.NewActivityUseCase(2, &activityRepo)
		activityDomainSample := []activities.Domain{
			{
				ID: 1,
				Name: "name",
				IsOutdoor: true,
				RecommendedWeather: "Clear",
				InterestID: 1,
			},
		}
		activityRepo.On("FindActivitiesByInterest", 
		mock.Anything, mock.AnythingOfType("[]int"), mock.AnythingOfType("bool")).Return(activityDomainSample)
		res := activityUC.GetActivitiesByInterest(ctx, []int{1,2}, true)
		
		assert.NotEmpty(t, res)
		assert.Equal(t, res, activityDomainSample)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		var activityRepo activityMock.Repository
		ctx := context.Background()
		activityUC := activities.NewActivityUseCase(2, &activityRepo)
		activityRepo.On("FindActivitiesByInterest", 
		mock.Anything, mock.AnythingOfType("[]int"), mock.AnythingOfType("bool")).Return([]activities.Domain{})
		res := activityUC.GetActivitiesByInterest(ctx, []int{1,2}, true)
		
		assert.Empty(t, res)
	})
}

func TestGetActivitiesInOut(t *testing.T){
	t.Run("test case 1, valid test", func(t *testing.T) {
		var activityRepo activityMock.Repository
		ctx := context.Background()
		activityUC := activities.NewActivityUseCase(2, &activityRepo)
		activityDomainSample := []activities.Domain{
			{
				ID: 1,
				Name: "name",
				IsOutdoor: true,
				RecommendedWeather: "Clear",
				InterestID: 1,
			},
		}
		activityRepo.On("FindActivitiesInOut", 
		mock.Anything, mock.AnythingOfType("bool")).Return(activityDomainSample)
		res := activityUC.GetActivitiesInOut(ctx, true)
		
		assert.NotEmpty(t, res)
		assert.Equal(t, res, activityDomainSample)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		var activityRepo activityMock.Repository
		ctx := context.Background()
		activityUC := activities.NewActivityUseCase(2, &activityRepo)
		activityRepo.On("FindActivitiesInOut", 
		mock.Anything, mock.AnythingOfType("bool")).Return([]activities.Domain{})
		res := activityUC.GetActivitiesInOut(ctx, true)
		
		assert.Empty(t, res)
	})
}