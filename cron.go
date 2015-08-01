package cron

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type Ticker struct {
	C    <-chan time.Time // The channel on which the ticks are delivered.
	c    chan time.Time   // Internal reference to C that publishes ticks.
	done chan struct{}    // The channel on which the stop signal is delivered.
	e    *cronexpr.Expression
}

func NewTicker(spec string) *Ticker {
	c := make(chan time.Time, 1)
	t := &Ticker{
		C:    c,
		c:    c,
		done: make(chan struct{}, 1),
		e:    cronexpr.MustParse(spec),
	}

	go t.start()
	return t
}

func (t *Ticker) start() {
	for {
		select {
		case <-time.After(t.e.Next(time.Now()).Sub(time.Now())):
			t.c <- time.Now()
		case <-t.done:
			break
		}
	}
}

func (t *Ticker) Stop() {
	t.done <- struct{}{}
	close(t.C)
}
