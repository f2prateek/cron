package main

import (
	"fmt"
	"os"

	"github.com/f2prateek/cron"
)

func main() {
	t := cron.New(os.Args[1])
	cmd := os.Args[2]
	args := os.Args[3:len(os.Args)]

	for {
		<-t.C
		fmt.Println("ticked", cmd, args)
	}
}
