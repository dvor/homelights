package main

import (
    "math"
    "time"
)

const kNormalizedSummerSolstice = 183
const kDaysInYear = 366

type TimeUtils interface {
    nextIterationDuration() time.Duration
    wakeUpTimeAt(date time.Time) time.Time
    sleepTimeAt(date time.Time) time.Time
}

func NewTimeUtils() TimeUtils {
    return TimeUtilsStruct{
        ts: NewTimeSource(),
    }
}

type TimeUtilsStruct struct {
    ts TimeSource
}

func (t TimeUtilsStruct) nextIterationDuration() time.Duration {
    now := t.ts.Now()

    wake := t.wakeUpTimeAt(now)
    sleep := t.sleepTimeAt(now)
    halfBS := sleep.Add(-30 * time.Minute)

    if now.Before(wake) {
        return wake.Sub(now)
    }
    if now.After(halfBS) && now.Before(sleep) {
        return sleep.Sub(now)
    }
    if now.After(sleep) {
        midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
        till := t.wakeUpTimeAt(midnight.Add(24 * time.Hour))
        return till.Sub(now)
    }

    return kDefaultUpdateInterval
}

func (t TimeUtilsStruct) wakeUpTimeAt(date time.Time) time.Time {
    midnight := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

    return midnight.Add(kWakeUpTime)
}

func (t TimeUtilsStruct) sleepTimeAt(date time.Time) time.Time {
    midnight := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

    day := date.YearDay()

    day = t.daysTillWinterSolstice(day)

    percent := float64(day) / kNormalizedSummerSolstice
    delta := float64(kSummerSleepTime - kWinterSleepTime) * percent

    ti := kWinterSleepTime + time.Duration(delta)
    res := midnight.Add(ti)

    return time.Date(res.Year(), res.Month(), res.Day(), res.Hour(), res.Minute(), res.Second(), 0, res.Location())
}

func (t TimeUtilsStruct) daysTillWinterSolstice(day int) int {
    // Normalize to start on winter solstice.
    day += 10

    if day > kNormalizedSummerSolstice {
        day = kDaysInYear - day
    }

    return int(math.Abs(float64(day)))
}
