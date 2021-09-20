package openweather

import "weatherit/usecases/weatherforecast"

type Weather struct {
	ID          int    `json:"id"`
	Name        string `json:"main"`
	Description string `json:"description"`
}

type WeatherForecast struct {
	DT      int64     `json:"dt"`
	Weather []Weather `json:"weather"`
}

type AllWeatherForecast struct {
	Current WeatherForecast   `json:"current"`
	Hourly  []WeatherForecast `json:"hourly"`
	Daily   []WeatherForecast `json:"daily"`
}

func (w *Weather) ToDomain() weatherforecast.Domain{
	return weatherforecast.Domain{
		ID: w.ID,
		Name: w.Name,
		Description: w.Description,
	}
}