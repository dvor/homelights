package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type Weather interface {
    weatherConditions() (int, time.Time, time.Time)
}

type WeatherStruct struct {
    config Config
}

func NewWeather(config Config) Weather {
    return WeatherStruct{config}
}

func (w WeatherStruct) weatherConditions() (int, time.Time, time.Time) {
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",
        w.config.Location.Latitude,
        w.config.Location.Longitude,
        w.config.Tokens.Weather,
    )
    res, err := http.Get(url)

    if err != nil {
        panic(NewError("Unable to get weather stats:", err))
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(NewError("Unable to read weather response:", err))
    }

    var dat map[string]interface{}
    err = json.Unmarshal(body, &dat)

    if err != nil {
        panic(NewError("Unable to parse weather json:", err))
    }

    cloudness := dat["clouds"].(map[string]interface{})["all"].(float64)
    sunrise := dat["sys"].(map[string]interface{})["sunrise"].(float64)
    sunset := dat["sys"].(map[string]interface{})["sunset"].(float64)

    sunriseTime := time.Unix(int64(sunrise), 0)
    sunsetTime := time.Unix(int64(sunset), 0)

    return int(cloudness), sunriseTime, sunsetTime
}
