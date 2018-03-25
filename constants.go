package main

import (
    "time"
)

const (
    kDefaultUpdateInterval = 30 * time.Minute

    kLightChangeDuration = 20 * time.Minute

    // Time to wake up, both summer and winter
    kWakeUpTime = 8 * time.Hour

    // Winter time to go to sleep, on 22 Dec
    kWinterSleepTime = 20 * time.Hour

    // Summer time to go to sleep, on 22 June
    kSummerSleepTime = 22 * time.Hour

    kCloudnessThreshold = 40
)
