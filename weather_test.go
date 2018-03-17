package main

import (
    "time"
)

type WeatherMock struct {
    cloudness int
    sunrise time.Time
    sunset time.Time
}

func (w WeatherMock) weatherConditions() (int, time.Time, time.Time) {
    return w.cloudness, w.sunrise, w.sunset
}
