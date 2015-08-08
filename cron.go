package cron

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type Ticker struct {
	C    <-chan time.Time // The channel on which the ticks are delivered.
	done chan struct{}    // The channel on which the stop signal is delivered.
	e    *cronexpr.Expression
}

func NewTicker(spec string) *Ticker {
	c := make(chan time.Time, 1)
	t := &Ticker{
		C:    c,
		done: make(chan struct{}, 1),
		e:    cronexpr.MustParse(spec),
	}

	go func() {
		for {
			select {
			case <-time.After(t.e.Next(time.Now()).Sub(time.Now())):
				c <- time.Now()
			case <-t.done:
				break
			}
		}
	}()

	return t
}

func (t *Ticker) Stop() {
	t.done <- struct{}{}
}
