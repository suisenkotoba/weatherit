package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const OpenWeatherOneApiUrl = "https://api.openweathermap.org/data/2.5/onecall?"

type OpenWeather struct {
	AppKey string
}

func (ow *OpenWeather) GetForecast(lat, long float64, excludes string) (res AllWeatherForecast) {
	url := OpenWeatherOneApiUrl + fmt.Sprintf("appid=%s?long=%f?lat=%f?exclude=%s", ow.AppKey, lat, long, excludes)
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

func (ow *OpenWeather) GetCurrentForecast(lat, long float64) Weather {
	excludes := "minutely,daily,hourly"
	res := ow.GetForecast(lat, long, excludes)
	return res.Current.Weather[0]
}

func (ow *OpenWeather) GetHourlyForecast(lat, long float64) []WeatherForecast {
	excludes := "minutely,daily"
	res := ow.GetForecast(lat, long, excludes)
	return res.Hourly
}

func (ow *OpenWeather) GetDailyForecast(lat, long float64) []WeatherForecast {
	excludes := "minutely,hourly"
	res := ow.GetForecast(lat, long, excludes)
	return res.Daily
}

func (ow *OpenWeather) GetTargetDTForecast(lat, long float64, dt1, dt2 int64, mode string) Weather {
	targetForecast := Weather{}
	forecasts := []WeatherForecast{}
	if mode == "hour" {
		forecasts = ow.GetHourlyForecast(lat, long)
	} else if mode == "day" {
		forecasts = ow.GetDailyForecast(lat, long)
	}
	for i := 0; i < len(forecasts); i++ {
		if forecasts[i].DT >= dt1 && forecasts[i].DT <= dt2 {
			return forecasts[i].Weather[0]
		}
	}
	return targetForecast
}