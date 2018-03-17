package main

import (
    "time"
)

type TimeSourceMock struct {
    t time.Time
}

func (t TimeSourceMock) Now() time.Time {
    return t.t
}
