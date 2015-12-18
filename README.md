cron
====

Cron library for golang.

Usage
=====
```go
// Create a ticker.
ticker := cron.New("40 12 * * * *")

// Wait for a tick. Run inside in a loop if you want to wait repeatedly.
t := <-time.After(next.Sub(time.Now()))

// Do some work.
doWork()

// Stop the ticker.
t.Stop()
```