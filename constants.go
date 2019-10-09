package main

import (
	"time"
)

const (
	kDefaultUpdateInterval = 20 * time.Minute

	// Interval to retry after panic.
	kPanicUpdateInterval = 1 * time.Minute

	kLightChangeDuration = 15 * time.Minute

	// Time to wake up, both summer and winter
	kWakeUpTime = 8 * time.Hour

	// Winter time to go to sleep, on 22 Dec
	kWinterSleepTime = 20 * time.Hour

	// Summer time to go to sleep, on 22 June
	kSummerSleepTime = 22 * time.Hour

	kCloudnessThreshold = 20

	// Delta from sunrise to turn lights on.
	kSunriseDelta = 1*time.Hour + 30*time.Minute

	// Delta to sunset to turn lights on.
	kSunsetDelta = 2 * time.Hour
)
