package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

type TimeUtilsMock struct {
    niDuration time.Duration
    wakeUpTime time.Time
    sleepTime time.Time
}

func (t TimeUtilsMock) nextIterationDuration() time.Duration {
    return t.niDuration
}

func (t TimeUtilsMock) wakeUpTimeAt(date time.Time) time.Time {
    return t.wakeUpTime
}

func (t TimeUtilsMock) sleepTimeAt(date time.Time) time.Time {
    return t.sleepTime
}

func TestNextIteractionDuration(t *testing.T) {
    assert := assert.New(t)

    // wake up 8:00:00, sleep 21:14:45
    const layout = "02.01.06 15:04:05 -0700"
    midnight, _ := time.Parse(layout, "14.04.18 00:00:00 +0100")

    ts := TimeSourceMock{t: midnight}
    tu := TimeUtilsStruct{ts: &ts}

    res := tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(3 * time.Hour + 27 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(8 * time.Hour)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(14 * time.Hour + 4 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(18 * time.Hour + 55 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(20 * time.Hour + 43 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(20 * time.Hour + 47 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 27 * time.Minute + 45 * time.Second)

    ts.t = midnight.Add(21 * time.Hour + 14 * time.Minute + 50 * time.Second)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)

    ts.t = midnight.Add(23 * time.Hour + 40 * time.Minute)
    res = tu.nextIterationDuration()
    assert.Equal(res, 30 * time.Minute)
}

func TestSleepTimeAt(t *testing.T) {
    assert := assert.New(t)
    tu := NewTimeUtils()

    layout := "02.01.06 15:04:05 -0700"
    var date time.Time

    date, _ = time.Parse(layout, "22.12.18 20:00:00 +0100")
    assert.Equal(date, tu.sleepTimeAt(date))

    date, _ = time.Parse(layout, "22.06.18 22:00:00 +0100")
    assert.Equal(date, tu.sleepTimeAt(date))

    date, _ = time.Parse(layout, "07.02.18 20:31:28 +0100")
    assert.Equal(date, tu.sleepTimeAt(date))

    date, _ = time.Parse(layout, "15.08.18 21:24:35 +0100")
    assert.Equal(date, tu.sleepTimeAt(date))

    date, _ = time.Parse(layout, "21.12.18 20:00:39 +0100")
    assert.Equal(date, tu.sleepTimeAt(date))
}
