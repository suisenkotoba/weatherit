package events_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"
	eventsMock "weatherit/mocks/events"
	weatherForecastMock "weatherit/mocks/weatherforecast"
	coordinate "weatherit/usecases/coordinates"
	events "weatherit/usecases/events"
	"weatherit/usecases/weatherforecast"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetAllUserEvents(t *testing.T) {
	t.Run("test case 1, valid  case (from & to)", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domainSample := []events.Domain{
			{
				ID:          1,
				UserID:      1,
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Title:       "title",
				Description: "desc",
				Address:     "address",
				GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
				EventChecklist: []events.Checklist{
					{
						ID:        1,
						Name:      "name",
						IsChecked: true,
					},
				},
			},
		}
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(domainSample, nil)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "2021-09-07", "2021-09-14", "")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("test case 2, valid  case (month)", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domainSample := []events.Domain{
			{
				ID:          1,
				UserID:      1,
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Title:       "title",
				Description: "desc",
				Address:     "address",
				GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
				EventChecklist: []events.Checklist{
					{
						ID:        1,
						Name:      "name",
						IsChecked: true,
					},
				},
			},
		}
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(domainSample, nil)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "", "", "2021-09")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("test case 3, valid  case (all)", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domainSample := []events.Domain{
			{
				ID:          1,
				UserID:      1,
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Title:       "title",
				Description: "desc",
				Address:     "address",
				GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
				EventChecklist: []events.Checklist{
					{
						ID:        1,
						Name:      "name",
						IsChecked: true,
					},
				},
			},
		}
		eventRepo.On("Find", mock.Anything, mock.AnythingOfType("int")).Return(domainSample, nil)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "", "", "")

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("test case 4, error raised (repo by date range)", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]events.Domain{}, errorRepo)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "2021-09-07", "2021-09-14", "")

		assert.NotNil(t, err)
		assert.Empty(t, res)
	})

	t.Run("test case 5, error raised (repo by month)", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]events.Domain{}, errorRepo)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "", "", "2021-09")

		assert.NotNil(t, err)
		assert.Empty(t, res)
	})

	t.Run("test case 6, error raised (repo all)", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("Find", mock.Anything, mock.AnythingOfType("int")).Return([]events.Domain{}, errorRepo)
		res, err := eventsUC.GetAllUserEvents(ctx, 1, "", "", "")

		assert.NotNil(t, err)
		assert.Empty(t, res)
	})
}

func TestGetAllUserEventsByDateRange(t *testing.T) {
	t.Run("test case 1, valid  case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domainSample := []events.Domain{
			{
				ID:          1,
				UserID:      1,
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Title:       "title",
				Description: "desc",
				Address:     "address",
				GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
				EventChecklist: []events.Checklist{
					{
						ID:        1,
						Name:      "name",
						IsChecked: true,
					},
				},
			},
		}
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(domainSample, nil)
		from, _ := time.Parse("2006-01-02", "2021-09-07")
		to, _ := time.Parse("2006-01-02", "2021-09-14")
		res, err := eventsUC.GetAllUserEventsByDateRange(ctx, 1, from, to)

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("FindByDate", mock.Anything, mock.AnythingOfType("int"),
			mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]events.Domain{}, errorRepo)
		from, _ := time.Parse("2006-01-02", "2021-09-07")
		to, _ := time.Parse("2006-01-02", "2021-09-14")
		res, err := eventsUC.GetAllUserEventsByDateRange(ctx, 1, from, to)

		assert.NotNil(t, err)
		assert.Empty(t, res)
	})
}
func TestGetAllEventByDateRange(t *testing.T) {
	t.Run("test case 1, valid  case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domainSample := []events.Domain{
			{
				ID:          1,
				UserID:      1,
				StartAt:     time.Now(),
				EndAt:       time.Now(),
				Title:       "title",
				Description: "desc",
				Address:     "address",
				GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
				EventChecklist: []events.Checklist{
					{
						ID:        1,
						Name:      "name",
						IsChecked: true,
					},
				},
			},
		}
		eventRepo.On("FindAllByDate", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return(domainSample, nil)
		from, _ := time.Parse("2006-01-02", "2021-09-07")
		to, _ := time.Parse("2006-01-02", "2021-09-14")
		res, err := eventsUC.GetAllEventByDateRange(ctx, from, to)

		assert.Nil(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("FindAllByDate", mock.Anything, mock.AnythingOfType("time.Time"), mock.AnythingOfType("time.Time")).Return([]events.Domain{}, errorRepo)
		from, _ := time.Parse("2006-01-02", "2021-09-07")
		to, _ := time.Parse("2006-01-02", "2021-09-14")
		res, err := eventsUC.GetAllEventByDateRange(ctx, from, to)

		assert.NotNil(t, err)
		assert.Empty(t, res)
	})
}
func TestScheduleEvent(t *testing.T) {
	t.Run("test case 1, valid case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domain := events.Domain{
			UserID:      1,
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Title:       "title",
			Description: "desc",
			Address:     "address",
			GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
			EventChecklist: []events.Checklist{
				{
					Name:      "name",
					IsChecked: true,
				},
			},
		}
		eventRepo.On("Store", mock.Anything, mock.AnythingOfType("*events.Domain")).Return(1, nil)
		eventId, err := eventsUC.ScheduleEvent(ctx, &domain)

		assert.Nil(t, err)
		assert.NotEqual(t, eventId, 0)

	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domain := events.Domain{
			UserID:      1,
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Title:       "title",
			Description: "desc",
			Address:     "address",
			GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
			EventChecklist: []events.Checklist{
				{
					Name:      "name",
					IsChecked: true,
				},
			},
		}
		eventRepo.On("Store", mock.Anything, mock.AnythingOfType("*events.Domain")).Return(0, errorRepo)
		eventId, err := eventsUC.ScheduleEvent(ctx, &domain)

		assert.NotNil(t, err)
		assert.Equal(t, eventId, 0)

	})
}
func TestCancelEvent(t *testing.T) {
	t.Run("test case 1, valid case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		err := eventsUC.CancelEvent(ctx, 1, 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		eventRepo.On("Delete", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errorRepo)
		err := eventsUC.CancelEvent(ctx, 1, 1)

		assert.NotNil(t, err)
	})
}
func TestUpdateEvent(t *testing.T) {
	t.Run("test case 1, valid case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domain := events.Domain{
			ID:          1,
			UserID:      1,
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Title:       "title",
			Description: "desc",
			Address:     "address",
			GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
			EventChecklist: []events.Checklist{
				{
					ID:        1,
					Name:      "name",
					IsChecked: true,
				},
			},
		}
		eventRepo.On("Update", mock.Anything, mock.AnythingOfType("*events.Domain")).Return(1, nil)
		err := eventsUC.UpdateEvent(ctx, &domain)

		assert.Nil(t, err)
	})

	t.Run("test case 2, error raised", func(t *testing.T) {
		errorRepo := errors.New("Error Repo")
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		ctx := context.Background()
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		domain := events.Domain{
			ID:          1,
			UserID:      1,
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Title:       "title",
			Description: "desc",
			Address:     "address",
			GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
			EventChecklist: []events.Checklist{
				{
					ID:        1,
					Name:      "name",
					IsChecked: true,
				},
			},
		}
		eventRepo.On("Update", mock.Anything, mock.AnythingOfType("*events.Domain")).Return(0, errorRepo)
		err := eventsUC.UpdateEvent(ctx, &domain)

		assert.NotNil(t, err)
	})
}
func TestForecastEvent(t *testing.T) {
	domain := events.Domain{
		ID:          1,
		UserID:      1,
		StartAt:     time.Now(),
		EndAt:       time.Now(),
		Title:       "title",
		Description: "desc",
		Address:     "address",
		GeoLoc:      coordinate.Coordinate{Lat: -1.43534, Long: 2.46433},
		EventChecklist: []events.Checklist{
			{
				ID:        1,
				Name:      "name",
				IsChecked: true,
			},
		},
	}
	forecast := weatherforecast.Domain{
		ID:          800,
		Name:        "name",
		Description: "desc",
	}
	t.Run("test case 1, valid case", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		weatherForecaster.On("GetTargetDTForecast", mock.AnythingOfType("float64"), mock.AnythingOfType("float64"),
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(forecast)
		dt := time.Now().Unix()
		res := eventsUC.ForecastEvent(domain, "hour", dt, dt + 60000)

		assert.NotEmpty(t, res)
	})

	t.Run("test case 2, invalid mode", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		dt := time.Now().Unix()
		res := eventsUC.ForecastEvent(domain, "minute", dt, dt + 60000)

		assert.Empty(t, res)
	})

	t.Run("test case 3, error repo", func(t *testing.T) {
		var (
			eventRepo         eventsMock.Repository
			weatherForecaster weatherForecastMock.Repository
		)
		eventsUC := events.NewEventUseCase(2, &eventRepo, &weatherForecaster)
		weatherForecaster.On("GetTargetDTForecast", mock.AnythingOfType("float64"), mock.AnythingOfType("float64"),
			mock.AnythingOfType("int64"), mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(weatherforecast.Domain{})
		dt := time.Now().Unix()
		res := eventsUC.ForecastEvent(domain, "minute", dt, dt + 60000)

		assert.Empty(t, res)
	})
	
}
