package ievio

import (
	"fmt"
	"syscall"
)

type InputEvent struct {
	Time  syscall.Timeval
	Type  EV_TYPE
	Code  Code
	Value int32
}

func (v *InputEvent) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", v.Time, v.Type, v.Code, v.Value)
}
