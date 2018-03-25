package main

import (
    "gopkg.in/gcfg.v1"
    "log"
)

var config = NewConfig()

type Config struct {
    Lights struct {
        Array []string
    }
    Tokens struct {
        Bridge string
        Weather string
    }
    Location struct {
        Latitude float64
        Longitude float64
    }
    Other struct {
        Notifierport string
    }
}

func NewConfig() Config {
    var c Config

    gcfg.ReadFileInto(&c, "config.gcfg")

    if len(c.Lights.Array) == 0 {
        log.Fatal("Please set lights array.")
    }

    if c.Tokens.Bridge == "" {
        log.Fatal("Please specify the weather token.")
    }

    if c.Tokens.Weather == "" {
        log.Fatal("Please specify the bridge api token.")
    }

    return c
}
