package main

import (
	"flag"
	"time"

	"github.com/gen2brain/beeep"
)

var units time.Duration = time.Second

func main() {
	interval := flag.Int("interval", 15, "Interval in seconds between notifications")
	flag.Parse()

	duration := time.Duration(*interval) * units

	for {
		err := beeep.Notify("Blink", "blink twice!", "logo.jpeg")
		if err != nil {
			panic(err)
		}
		time.Sleep(duration)
	}
}
