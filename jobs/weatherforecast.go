package jobs

import (
	"context"
	"math/rand"
	"time"
	"weatherit/usecases"
	"weatherit/usecases/alterplan"
	"weatherit/usecases/events"
)

func Forecast(uc usecases.UsecaseList, offset string) {
	niceWeatherIndex := map[string]int{
		"Thunderstorm": 200,
		"Drizzle":      300,
		"Rain":         500,
		"Snow":         600,
		"Atmosphere":   700,
		"Clear":        800,
		"Clouds":       800}

	ctx := context.TODO()

	events := []events.Domain{}

	today := time.Now()
	from := time.Date(today.Year(), today.Month(), today.Day(), today.Hour(), today.Minute(), 0, 0, today.Location())
	if offset == "H1" {
		from = from.Add(60 * 60 * 1000000000)
		to := from.Add(60 * 1000000000)
		events, _ = uc.Event.GetAllEventByDateRange(ctx, from, to)

	} else if offset == "H6" {
		from = from.Add(6 * 60 * 60 * 1000000000)
		to := from.Add(60 * 1000000000)
		events, _ = uc.Event.GetAllEventByDateRange(ctx, from, to)

	} else if offset == "D1" {
		from = from.AddDate(0, 0, 1)
		to := from.Add(60 * 1000000000)
		events, _ = uc.Event.GetAllEventByDateRange(ctx, from, to)
	}

	for i := 0; i < len(events); i++ {
		alterPlan := uc.AlterPlan.GetEventAlterPlan(ctx, events[i].ID)
		if alterPlan.ID == 0 {
			alterPlan = alterplan.Domain{
				EventID: events[i].ID,
			}
			alterPlan.ID, _ = uc.AlterPlan.MakeEventAlterPlan(ctx, &alterPlan)
		}
		dt1 := events[i].StartAt.Unix()
		dt2 := events[i].EndAt.Unix()
		if offset == "H1" {
			weather := uc.Event.ForecastEvent(events[i], "hour", dt1, dt2)
			alterPlan.WeatherForecastH1 = weather.Name
			if niceWeatherIndex[weather.Name] < 700 {
				interests := uc.UserInterest.GetUserInterestIDs(ctx, events[i].UserID)
				activities := uc.Activity.GetActivitiesByInterest(ctx, interests, true)
				randomIndex := rand.Intn(len(activities))
				alterPlan.ActivityID = activities[randomIndex].ID
				//	TODO send via pusher
			}

		} else if offset == "H6" {
			weather := uc.Event.ForecastEvent(events[i], "hour", dt1, dt2)
			alterPlan.WeatherForecastH6 = weather.Name
		} else if offset == "D1" {
			weather := uc.Event.ForecastEvent(events[i], "day", dt1, dt2)
			alterPlan.WeatherForecastD1 = weather.Name
		}
		uc.AlterPlan.UpdateEventAlterPlan(ctx, &alterPlan)
	}
}
