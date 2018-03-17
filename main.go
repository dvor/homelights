package main

import (
    "github.com/collinux/gohue"
    "github.com/sevlyar/go-daemon"
    "gopkg.in/gcfg.v1"
    "log"
    "net/http"
    "os"
    "time"
)

type Config struct {
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
var config Config

const serverPath = "./html"

var actionManager = NewActionManager()
var timeUtils = NewTimeUtils()
var notifier = NewStatusNotifier(serverPath)

func main() {
    gcfg.ReadFileInto(&config, "config.gcfg")

    if config.Tokens.Bridge == "" {
        log.Fatal("Please specify the weather token.")
    }

    if config.Tokens.Weather == "" {
        log.Fatal("Please specify the bridge api token.")
    }

    cntxt := daemonContext()
    child, err := cntxt.Reborn()
    if err != nil {
        log.Fatal("Unable to run: ", err)
    }
    if child != nil {
        return
    }

    log.Print("- - - - - - - - - - - - - - -")
    log.Print("daemon started")
    defer log.Print("daemon exited")
    defer cntxt.Release()

    go runLoop()

    os.Mkdir(serverPath, os.ModePerm)
    http.Handle("/homelights/", http.StripPrefix("/homelights/", http.FileServer(http.Dir(serverPath))))
    http.ListenAndServe(config.Other.Notifierport, nil)
}

func daemonContext() *daemon.Context {
    return &daemon.Context{
        PidFileName: "pid",
        PidFilePerm: 0644,
        LogFileName: "log",
        LogFilePerm: 0640,
        WorkDir:     "./",
        Umask:       027,
        Args:        []string{"homelights daemon"},
    }
}

func runLoop() {
    for {
        iteration()

        d := timeUtils.nextIterationDuration()
        log.Print("Sleeping for", d)
        time.Sleep(d)
    }
}

func iteration() {
    log.Print("Updating...")

    defer notifier.update()
    notifier.reset()
    notifier.append("Running")
    notifier.append("")
    t := NewTimeSource().Now()
    notifier.append(t.Format("15:04:05 02/01/06"))
    notifier.append("")

    bridgesOnNetwork, err := hue.FindBridges()

    if err != nil || len(bridgesOnNetwork) == 0 {
        log.Print("Cannot find bridge, err", err)
        notifier.append("No bridge")
        time.Sleep(5 * time.Minute)
        return
    }

    bridge := bridgesOnNetwork[0]
    bridge.Login(config.Tokens.Bridge)

    light1, err1 := bridge.GetLightByName("Bird lamp 1")
    light2, err2 := bridge.GetLightByName("Bird lamp 2")

    if err1 != nil || err2 != nil {
        log.Print("Cannot find lights")
        log.Print("err1", err1)
        log.Print("err2", err2)
        notifier.append("No lights")
        time.Sleep(5 * time.Minute)
        return
    }

    a := actionManager.currentAction()
    switch a {
    case ActionOn:
        log.Print("Lights On")
        notifier.append("On")
        light1.On()
        light2.On()

        light1.SetBrightness(100)
        light2.SetBrightness(100)
    case ActionOff:
        notifier.append("Off")
        log.Print("Lights Off")
        light1.Off()
        light2.Off()
    }
}
