package cron

import (
	"time"

	"github.com/f2prateek/clock"
	"github.com/gorhill/cronexpr"
)

var c = clock.Default()

type Ticker struct {
	C    <-chan time.Time // The channel on which the ticks are delivered.
	done chan struct{}    // The channel on which the stop signal is delivered.
	expr *cronexpr.Expression
}

// Must is a helper that wraps a call to a function returning (*Ticker, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as `var t = cron.Must(cron.Parse("spec"))`
func Must(t *Ticker, err error) *Ticker {
	if err != nil {
		panic(err)
	}
	return t
}

// Parse returns a new Ticker containing a channel that will send the time
// with a period specified by the spec argument. Stop the ticker to release
// associated resources.
func Parse(spec string) (*Ticker, error) {
	expr, err := cronexpr.Parse(spec)
	if err != nil {
		return nil, err
	}
	tickerC := make(chan time.Time, 1)
	ticker := &Ticker{
		C:    tickerC,
		done: make(chan struct{}, 1),
		expr: expr,
	}

	go func() {
		for {
			next := ticker.expr.Next(c.Now())
			select {
			case <-time.After(next.Sub(c.Now())):
				tickerC <- c.Now()
			case <-ticker.done:
				break
			}
		}
	}()

	return ticker, nil
}

// Stop turns off a ticker. After Stop, no more ticks will be sent. Stop does
// not close the channel, to prevent a read from the channel succeeding
// incorrectly.
func (t *Ticker) Stop() {
	t.done <- struct{}{}
}
