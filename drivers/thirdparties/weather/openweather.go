package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weatherit/usecases/weatherforecast"
)

const OpenWeatherOneApiUrl = "https://api.openweathermap.org/data/2.5/onecall?"

type openWeather struct {
	AppKey string
}

func NewWeatherForecaster(appKey string) weatherforecast.Repository {
	return &openWeather{
		AppKey: appKey,
	}
}

func (ow *openWeather) GetForecast(lat, long float64, excludes string) (res AllWeatherForecast) {
	url := OpenWeatherOneApiUrl + fmt.Sprintf("appid=%s&lon=%f&lat=%f&exclude=%s", ow.AppKey, long, lat, excludes)
	response, err := http.Get(url)
	if err != nil {
		return res
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res
	}
	return

}

func (ow *openWeather) GetCurrentForecast(lat, long float64) weatherforecast.Domain {
	excludes := "minutely,daily,hourly"
	res := ow.GetForecast(lat, long, excludes)
	return res.Current.Weather[0].ToDomain()
}

func (ow *openWeather) GetHourlyForecast(lat, long float64) []WeatherForecast {
	excludes := "minutely,daily"
	res := ow.GetForecast(lat, long, excludes)
	return res.Hourly
}

func (ow *openWeather) GetDailyForecast(lat, long float64) []WeatherForecast {
	excludes := "minutely,hourly"
	res := ow.GetForecast(lat, long, excludes)
	return res.Daily
}

func (ow *openWeather) GetTargetDTForecast(lat, long float64, dt1, dt2 int64, mode string) weatherforecast.Domain {
	targetForecast := Weather{}
	forecasts := []WeatherForecast{}
	if mode == "hour" {
		forecasts = ow.GetHourlyForecast(lat, long)
	} else if mode == "day" {
		forecasts = ow.GetDailyForecast(lat, long)
	}
	for i := 0; i < len(forecasts); i++ {
		if forecasts[i].DT >= dt1 && forecasts[i].DT <= dt2 {
			return forecasts[i].Weather[0].ToDomain()
		}
	}
	return targetForecast.ToDomain()
}
