package main

import (
    "log"
    "time"
)

type Action int

const (
    ActionOff Action = iota
    ActionOn
)

type ActionManager struct {
    notifier *StatusNotifier
    ts TimeSource
    tu TimeUtils
    wt Weather
}

func NewActionManager(config Config, notifier *StatusNotifier) ActionManager {
    return ActionManager{
        notifier: notifier,
        ts: NewTimeSource(),
        tu: NewTimeUtils(),
        wt: NewWeather(config),
    }
}

func (a ActionManager) currentAction() Action {
    now := a.ts.Now()
    wakeUp := a.tu.wakeUpTimeAt(now)
    sleep := a.tu.sleepTimeAt(now)
    cloudness, sunrise, sunset := a.wt.weatherConditions()

    a.notifier.append(cloudness)

    log.Print("Status:")
    log.Print("     Wake up:", wakeUp)
    log.Print("     Sleep:  ", sleep)
    log.Print("     Clouds: ", cloudness)
    log.Print("     Sunrise:", sunrise)
    log.Print("     Sunset: ", sunset)

    if now.Before(wakeUp) || now.After(sleep) {
        return ActionOff
    }

    if now.Before(sunrise.Add(1 * time.Hour)) {
        return ActionOn
    }

    if now.Before(sunset.Add(-1 * time.Hour)) {
        if cloudness > kCloudnessThreshold {
            return ActionOn
        }
        return ActionOff
    }

    // after sunset-1h, but before sleep
    return ActionOn
}
