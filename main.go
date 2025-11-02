package main

import (
	"flag"
	"time"

	"github.com/gen2brain/beeep"
)

var units time.Duration = time.Second

func main() {
	interval := flag.Int("interval", 15, "Interval in seconds between notifications")
	minutes := flag.Bool("minutes", false, "Interval units are in seconds (default) or minutes if true")
	flag.Parse()

	if *minutes {
		units = time.Minute
	}

	tickerBreak := time.NewTicker(20 * time.Minute)
	tickerBlink := time.NewTicker(time.Duration(*interval) * (units))
	defer tickerBreak.Stop()
	defer tickerBlink.Stop()

	for {
		select {
		case <-tickerBreak.C:
			err := beeep.Notify("Break", "Look far away for 20sec!", "")
			if err != nil {
				panic(err)
			}
			// drain the blink ticker during the break
			select {
			case <-tickerBlink.C:
			default:
			}
		case <-tickerBlink.C:
			err := beeep.Notify("Blink", "blink twice!", "")
			if err != nil {
				panic(err)
			}
		}
	}
}
