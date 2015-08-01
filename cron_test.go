package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	ticker := NewTicker("3 13 * * * *")

	select {
	case t := <-ticker.C:
		fmt.Println("ticked", t)
	case <-time.After(1 * time.Minute):
		fmt.Println("timed out")
	}
}
