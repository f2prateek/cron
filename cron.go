package cron

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type Ticker struct {
	C    <-chan time.Time // The channel on which the ticks are delivered.
	done chan struct{}    // The channel on which the stop signal is delivered.
	expr *cronexpr.Expression
}

// NewTicker returns a new Ticker containing a channel that will send the time
// with a period specified by the spec argument. Stop the ticker to release
// associated resources.
func New(spec string) (*Ticker, error) {
	expr, err := cronexpr.Parse(spec)
	if err != nil {
		return nil, err
	}
	c := make(chan time.Time, 1)
	t := &Ticker{
		C:    c,
		done: make(chan struct{}, 1),
		expr: expr,
	}

	go func() {
		for {
			next := t.expr.Next(time.Now())
			select {
			case <-time.After(next.Sub(time.Now())):
				c <- time.Now()
			case <-t.done:
				break
			}
		}
	}()

	return t, nil
}

// Stop turns off a ticker. After Stop, no more ticks will be sent. Stop does
// not close the channel, to prevent a read from the channel succeeding
// incorrectly.
func (t *Ticker) Stop() {
	t.done <- struct{}{}
}
