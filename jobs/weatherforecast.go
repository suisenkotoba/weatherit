package jobs

import (
	"context"
	"fmt"
	"time"
	"weatherit/usecases"
	"weatherit/usecases/alterplan"
	"weatherit/usecases/events"
)

func Forecast(uc usecases.UsecaseList, offset string) {
	// get all event within range
	ctx := context.TODO()
	data, err := uc.Interest.GetAvailableInterests(ctx)
	fmt.Println(err)
	fmt.Println(data)
	// ==

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
		}
		if offset == "H1" {
			interests := uc.UserInterest.GetUserInterestIDs(ctx, events[i].UserID)
			activities := uc.Activity.GetActivitiesByInterest(ctx, interests)
			//	random pick activity
			//	create alter plan for event
			//	send via pusher
		}
	}
}
