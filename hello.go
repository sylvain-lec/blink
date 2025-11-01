package main

import (
	"time"

	"github.com/gen2brain/beeep"
)

var (
	count int           = 15
	units time.Duration = time.Second
)

func main() {
	duration := time.Duration(count) * units

	for {
		err := beeep.Notify("Blink", "blink twice!", "logo.jpeg")
		if err != nil {
			panic(err)
		}
		time.Sleep(duration)
	}
}
