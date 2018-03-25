package main

import (
    "io/ioutil"
    "fmt"
    "log"
)

type Notifier struct {
    path string
    msg string
}

func NewNotifier(path string) Notifier {
    s := Notifier{path: path}
    s.reset()
    return s
}

func (s *Notifier) reset() {
    s.msg = ""
}

func (s *Notifier) append(v ...interface{}) {
    str := fmt.Sprint(v...)
    s.msg = s.msg + str + "\n"
}

func (s *Notifier) update() {
    t := NewTimeSource().Now()
    msg := t.Format("15:04:05 02/01/06") + "\n\n" + s.msg

    b := []byte(msg)
    p := s.path + "/status"

    err := ioutil.WriteFile(p, b, 0644)

    if err != nil {
        log.Print("Cannot write file", p)
    }
}

