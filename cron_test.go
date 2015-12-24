package cron_test

import (
	"testing"
	"time"

	"github.com/f2prateek/cron"
)

func TestCron(t *testing.T) {
	// Schedule a cron ticker for every 5 seconds.
	ticker := cron.New("0/5 * * * * * *")
	defer ticker.Stop()

	// Wait for the tick to be delivered with 10 seconds.
	select {
	case <-ticker.C:
	case <-time.After(10 * time.Second):
		t.Error("cron failed to deliver tick in time")
	}
}
