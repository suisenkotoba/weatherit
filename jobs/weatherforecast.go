package jobs

import (
	"context"
	"fmt"
	"weatherit/usecases"
)

func Forecast(uc usecases.UsecaseList, offset string) {
	// get all event within
	ctx := context.TODO()
	data, err := uc.Interest.GetAvailableInterests(ctx)
	fmt.Println(err)
	fmt.Println(data)

	// events := []events.Domain{}

	if offset == "H1" {
		// next  hour second 00 + 1 minute
	} else if offset == "H6" {
		// next 6 hours second 00 + 1 minute

	} else if offset == "D1" {
		// tomorrow second 00 + 1 minute
	}

	// for every event, check weather
	// if weather not suitable for outdoor,
	// 	get user interest
	//	get list activities from collected interests
	//	random pick activity
	//	create alter plan for event
	//	send via pusher
}
