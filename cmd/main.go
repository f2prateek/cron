package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/f2prateek/cron"
)

func main() {
	t, err := cron.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[2]
	args := os.Args[3:len(os.Args)]

	for {
		<-t.C

		cmd := exec.Command(cmd, args...)
		out, err := cmd.CombinedOutput()
		fmt.Printf(string(out))
		if err != nil {
			log.Fatal(err)
		}
	}
}
