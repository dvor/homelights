package main

import (
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
	notifier := NewNotifier(serverPath)

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

func runLoop(config Config, notifier *Notifier) {
	for {
		safeIteration(config, notifier)
	}
}

func safeIteration(config Config, notifier *Notifier) {
	defer func() {
		if err := recover(); err != nil {
			notifier.reset()
			notifier.append("Panic")
			notifier.append(err)
			notifier.update()

			log.Print("Panic! Iteration failed with error: ", err)
			log.Print("Sleeping for ", kPanicUpdateInterval)
			time.Sleep(kPanicUpdateInterval)
		}
	}()

	iteration(config, notifier)
}

func iteration(config Config, notifier *Notifier) {
	log.Print("Updating...")

	notifier.reset()
	notifier.append("Running")

	bridge := FindBridge(config)
	room := bridge.room()

	a := NewActionManager(config, notifier).currentAction()
	notifier.append()

	switch a {
	case ActionOn:
		notifier.append("On")
		go room.changeTo(true)
	case ActionOff:
		notifier.append("Off")
		go room.changeTo(false)
	}

	notifier.update()

	d := NewTimeUtils().nextIterationDuration()
	log.Print("Sleeping for ", d)
	time.Sleep(d)
}
