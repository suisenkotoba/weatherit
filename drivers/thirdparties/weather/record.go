package openweather

type Weather struct {
	ID          int    `json:"id"`
	Name        string `json:"main"`
	Description string `json:"description"`
}

type WeatherForecast struct {
	DT      int       `json:"dt"`
	Weather []Weather `json:"weather"`
}

type AllWeatherForecast struct {
	Current WeatherForecast   `json:"current"`
	Hourly  []WeatherForecast `json:"hourly"`
	Daily   []WeatherForecast `json:"daily"`
}
