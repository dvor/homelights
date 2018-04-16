package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestCurrentAction(t *testing.T) {
	assert := assert.New(t)
	log.SetOutput(ioutil.Discard)

	const layout = "02.01.06 15:04:05 -0700"
	midnight, _ := time.Parse(layout, "14.04.18 00:00:00 +0100")

	at := func(str string) time.Time {
		d, _ := time.ParseDuration(str)
		return midnight.Add(d)
	}

	nt := Notifier{}
	ts := TimeSourceMock{midnight}
	tu := TimeUtilsMock{0, midnight, midnight}
	wt := WeatherMock{0, midnight, midnight}
	am := ActionManager{&nt, &ts, &tu, &wt}

	tu.wakeUpTime = at("8h")
	tu.sleepTime = at("20h")
	wt.cloudness = 30
	wt.sunrise = at("6h")
	wt.sunset = at("22h")

	ts.t = at("0h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("4h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("7h35m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("8h35m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("10h20m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("11h11m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("14h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("16h42m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("17h13m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("19h13m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("20h10m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("23h20m")
	assert.Equal(ActionOff, am.currentAction())

	tu.wakeUpTime = at("8h")
	tu.sleepTime = at("20h")
	wt.cloudness = 50
	wt.sunrise = at("6h")
	wt.sunset = at("22h")

	ts.t = at("0h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("4h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("7h35m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("8h35m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("10h20m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("11h11m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("14h")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("16h42m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("17h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("19h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("20h10m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("23h20m")
	assert.Equal(ActionOff, am.currentAction())

	tu.wakeUpTime = at("8h")
	tu.sleepTime = at("20h")
	wt.cloudness = 30
	wt.sunrise = at("10h")
	wt.sunset = at("18h")

	ts.t = at("0h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("4h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("7h35m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("8h35m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("10h20m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("11h11m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("14h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("16h42m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("17h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("19h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("20h10m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("23h20m")
	assert.Equal(ActionOff, am.currentAction())

	tu.wakeUpTime = at("8h")
	tu.sleepTime = at("20h")
	wt.cloudness = 50
	wt.sunrise = at("10h")
	wt.sunset = at("18h")

	ts.t = at("0h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("4h")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("7h35m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("8h35m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("10h20m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("11h11m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("14h")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("16h42m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("17h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("19h13m")
	assert.Equal(ActionOn, am.currentAction())

	ts.t = at("20h10m")
	assert.Equal(ActionOff, am.currentAction())

	ts.t = at("23h20m")
	assert.Equal(ActionOff, am.currentAction())
}
