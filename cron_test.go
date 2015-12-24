package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	now := time.Now()
	hour, minute := now.Hour(), now.Minute()+1

	// Schedule cron ticker for the next minute.
	ticker := New(fmt.Sprintf("%d %d * * * *", minute, hour))

	// Wait for the tick for a minute.
	select {
	case <-ticker.C:
		fmt.Println("ticked")
	case <-time.After(1 * time.Minute):
		t.Error("cron failed to deliver tick in time")
	}

	ticker.Stop()
}
