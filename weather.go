package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type Weather interface {
    weatherConditions() (int, time.Time, time.Time)
}

type WeatherStruct struct {
}

func NewWeather() Weather {
    return WeatherStruct{}
}

func (w WeatherStruct) weatherConditions() (int, time.Time, time.Time) {
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",
        config.Location.Latitude,
        config.Location.Longitude,
        config.Tokens.Weather,
    )
    res, err := http.Get(url)

    if err != nil {
        log.Fatal("Unable to get weather stats: ", err)
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatal("Unable to read weather response: ", err)
    }

    var dat map[string]interface{}
    err = json.Unmarshal(body, &dat)

    if err != nil {
        log.Fatal("Unable to parse weather json: ", err)
    }

    cloudness := dat["clouds"].(map[string]interface{})["all"].(float64)
    sunrise := dat["sys"].(map[string]interface{})["sunrise"].(float64)
    sunset := dat["sys"].(map[string]interface{})["sunset"].(float64)

    sunriseTime := time.Unix(int64(sunrise), 0)
    sunsetTime := time.Unix(int64(sunset), 0)

    return int(cloudness), sunriseTime, sunsetTime
}