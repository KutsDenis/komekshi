package scheduler

import (
	"time"
)

type Scheduler interface {
	EachTime(function func(), timer Timer)
}

type Timer struct {
	Min uint16
}

func (t Timer) EachTime(function func()) {
	for {
		function()
		time.Sleep(time.Duration(t.Min) * time.Minute)
	}
}
