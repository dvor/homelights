package main

import (
    "github.com/collinux/gohue"
)

type Bridge struct {
    br hue.Bridge
    config Config
}

func FindBridge(config Config) Bridge {
    bridgesOnNetwork, err := hue.FindBridges()

    if err != nil {
        panic(NewError("Cannot find bridge:", err))
    }
    if len(bridgesOnNetwork) == 0 {
        panic(NewError("No bridges found"))
    }

    br := bridgesOnNetwork[0]
    err = br.Login(config.Tokens.Bridge)

    if err != nil {
        panic(NewError("Cannot login into bridge:", err))
    }

    return Bridge{br: br, config: config}
}

func (b *Bridge) room() Room {
    var lights []hue.Light

    for _, name := range b.config.Lights.Array {
        l, err := b.br.GetLightByName(name)

        if err != nil {
            panic(NewError("Cannot find light:", err))
        }

        lights = append(lights, l)
    }

    return NewRoom(lights, b.config)
}
