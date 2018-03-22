package main

import (
    "github.com/collinux/gohue"
    "github.com/sevlyar/go-daemon"
    "log"
    "net/http"
    "os"
    "time"
)


func main() {
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

    const serverPath = "./html"
    config := NewConfig()
    notifier := NewStatusNotifier(serverPath)

    go runLoop(config, &notifier)

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

func runLoop(config Config, notifier *StatusNotifier) {
    actionManager := NewActionManager(config, notifier)
    timeUtils := NewTimeUtils()

    for {
        safeIteration(actionManager, timeUtils, notifier)
    }
}

func safeIteration(actionManager ActionManager,
                   timeUtils TimeUtils,
                   notifier *StatusNotifier) {

    defer func() {
        if err := recover(); err != nil {
            notifier.reset()
            notifier.append("Panic")
            notifier.append(err)
            notifier.update()

            log.Print("Panic! Iteration failed with error: ", err)
            log.Print("Sleeping for 5 minutes")

            time.Sleep(5 * time.Minute)
        }
    }()

    iteration(actionManager, timeUtils, notifier)
}

func iteration(actionManager ActionManager,
               timeUtils TimeUtils,
               notifier *StatusNotifier) {
    log.Print("Updating...")

    notifier.reset()
    notifier.append("Running")

    bridgesOnNetwork, err := hue.FindBridges()

    if err != nil || len(bridgesOnNetwork) == 0 {
        panic(NewError("Cannot find bridge:", err))
    }

    bridge := bridgesOnNetwork[0]
    bridge.Login(config.Tokens.Bridge)

    light1, err1 := bridge.GetLightByName("Bird lamp 1")
    light2, err2 := bridge.GetLightByName("Bird lamp 2")

    if err1 != nil || err2 != nil {
        panic(NewError("Cannot find lights:", err1, err2))
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

    notifier.update()

    d := timeUtils.nextIterationDuration()
    log.Print("Sleeping for", d)
    time.Sleep(d)
}
