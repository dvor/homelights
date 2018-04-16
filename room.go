package main

import (
	"github.com/collinux/gohue"
	"log"
	"time"
)

type Room struct {
	lights []hue.Light
	config Config
}

func NewRoom(lights []hue.Light, config Config) Room {
	if len(lights) == 0 {
		panic(NewError("There should be at least one light"))
	}

	return Room{lights, config}
}

func (r *Room) isOn() bool {
	return r.lights[0].State.On
}

func (r *Room) changeTo(state bool) {
	log.Print("Room changeTo ", state)

	if r.isOn() == state {
		log.Print("Room is already ", state, ", exiting")
		return
	}

	if state {
		for _, l := range r.lights {
			l.SetBrightness(1)
			l.On()
		}
	}

	for i := 1; i <= 100; i++ {
		percent := i
		if !state {
			percent = 101 - i
		}

		log.Print("Room updating brightness ", percent)
		for _, l := range r.lights {
			l.SetBrightness(percent)
		}

		time.Sleep(kLightChangeDuration / 100)
	}

	if !state {
		for _, l := range r.lights {
			l.Off()
		}
	}
}
