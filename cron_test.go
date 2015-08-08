package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ticker := NewTicker("27 12 * * * *")

	select {
	case t := <-ticker.C:
		fmt.Println("ticked", t)
	case <-time.After(1 * time.Minute):
		t.Error("cron failed to deliver tick in time")
	}
}
