package main

import (
    "io/ioutil"
    "fmt"
    "log"
)

type StatusNotifier struct {
    path string
    msg string
}

func NewStatusNotifier(path string) StatusNotifier {
    s := StatusNotifier{path: path}
    s.reset()
    return s
}

func (s *StatusNotifier) reset() {
    s.msg = ""
}

func (s *StatusNotifier) append(v ...interface{}) {
    str := fmt.Sprint(v...)
    s.msg = s.msg + str + "\n"
}

func (s *StatusNotifier) update() {
    b := []byte(s.msg)
    p := s.path + "/status"

    err := ioutil.WriteFile(p, b, 0644)

    if err != nil {
        log.Print("Cannot write file", p)
    }
}

