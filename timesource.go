package main

import (
	"time"
)

type TimeSource interface {
	Now() time.Time
}

type TimeSourceStruct struct {
}

func NewTimeSource() TimeSource {
	return TimeSourceStruct{}
}

func (t TimeSourceStruct) Now() time.Time {
	return time.Now()
}
