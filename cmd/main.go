package main

import (
	"fmt"
	"log"
	"os"

	"github.com/f2prateek/cron"
)

func main() {
	t, err := cron.New(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[2]
	args := os.Args[3:len(os.Args)]

	for {
		<-t.C
		fmt.Println("ticked", cmd, args)
	}
}
