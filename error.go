package main

import "fmt"

type Error struct {
    str string
}

func NewError(v ...interface{}) Error {
    str := fmt.Sprint(v...)
    return Error{str}
}

func (e Error) Error() string {
    return e.str
}
