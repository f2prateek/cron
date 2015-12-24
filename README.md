cron
====

Cron library for golang.

Usage
=====
```go
// Create a ticker.
ticker := cron.Must(cron.Parse("0/5 * * * * * *"))

// Wait for a tick. Run inside in a loop if you want to wait repeatedly.
t := <-ticker.C

// Do some work.
doWork()

// Stop the ticker.
t.Stop()
```
