package alterplan_test

import (
	"context"
	"errors"
	"os"
	"testing"
	alterplanMock "weatherit/mocks/alterplan"
	alterplan "weatherit/usecases/alterplan"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetEventAlterPlan(t *testing.T) {
	t.Run("test case 1, valid  case", func(t *testing.T) {
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		alterplanDomainSample := alterplan.Domain{
			ID: 1,
			EventID: 1,
			ActivityID: 1,
			WeatherForecastH1: "forecast1",
			WeatherForecastH6: "forecast2",
			WeatherForecastD1: "forecast3",
			IsTaken: false,
		}
		alterplanRepo.On("FindByEventID", mock.Anything, mock.AnythingOfType("int")).Return(alterplanDomainSample).Once()
		res := alterplanUC.GetEventAlterPlan(ctx, 1)

		assert.NotEmpty(t, res)
		assert.Equal(t, alterplanDomainSample, res)
	})

	t.Run("test case 2, error raised/empty result", func(t *testing.T) {
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		
		alterplanRepo.On("FindByEventID", mock.Anything, mock.AnythingOfType("int")).Return(alterplan.Domain{}).Once()
		res := alterplanUC.GetEventAlterPlan(ctx, 1)

		assert.Empty(t, res)
	})

}

func TestMakeEventAlterPlan(t *testing.T) {
	t.Run("test case 1, valid case", func(t *testing.T) {
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		alterplanDomainSample := alterplan.Domain{
			EventID: 1,
			ActivityID: 1,
			WeatherForecastH1: "forecast1",
		}
		alterplanRepo.On("Store", mock.Anything, mock.AnythingOfType("*alterplan.Domain")).Return(1, nil)
		alterplanID, err := alterplanUC.MakeEventAlterPlan(ctx, &alterplanDomainSample)

		assert.Nil(t, err)
		assert.NotEqual(t, alterplanID, 0)

	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		alterplanDomainSample := alterplan.Domain{
			EventID: 1,
			ActivityID: 1,
			WeatherForecastH1: "forecast1",
		}
		alterplanRepo.On("Store", mock.Anything, mock.AnythingOfType("*alterplan.Domain")).Return(0, errorRepo)
		alterplanID, err := alterplanUC.MakeEventAlterPlan(ctx, &alterplanDomainSample)

		assert.NotEmpty(t, err)
		assert.Equal(t, alterplanID, 0)

	})
}

func TestUpdateEventAlterPlan(t *testing.T) {
	t.Run("test case 1, valid case", func(t *testing.T) {
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		alterplanDomainSample := alterplan.Domain{
			ID: 1,
			EventID: 1,
			ActivityID: 1,
			WeatherForecastH1: "forecast1",
			WeatherForecastH6: "forecast2",
			WeatherForecastD1: "forecast3",
			IsTaken: false,
		}
		alterplanRepo.On("Update", mock.Anything, mock.AnythingOfType("*alterplan.Domain")).Return(nil)
		err := alterplanUC.UpdateEventAlterPlan(ctx, &alterplanDomainSample)

		assert.Nil(t, err)

	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var alterplanRepo alterplanMock.Repository
		alterplanUC := alterplan.NewAlterPlanUseCase(2, &alterplanRepo)
		ctx := context.Background()
		alterplanDomainSample := alterplan.Domain{
			ID: 1,
			EventID: 1,
			ActivityID: 1,
			WeatherForecastH1: "forecast1",
			WeatherForecastH6: "forecast2",
			WeatherForecastD1: "forecast3",
			IsTaken: false,
		}
		alterplanRepo.On("Update", mock.Anything, mock.AnythingOfType("*alterplan.Domain")).Return(errorRepo)
		err := alterplanUC.UpdateEventAlterPlan(ctx, &alterplanDomainSample)

		assert.NotEmpty(t, err)

	})
}
